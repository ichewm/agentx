package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type benchReport struct {
	Passed    int           `json:"passed"`
	Failed    int           `json:"failed"`
	Scenarios []benchResult `json:"scenarios"`
}

type benchResult struct {
	Name     string   `json:"name"`
	Passed   bool     `json:"passed"`
	Expected string   `json:"expected"`
	Findings []string `json:"findings,omitempty"`
	Error    string   `json:"error,omitempty"`
}

type gateResult struct {
	Capability string   `json:"capability"`
	Target     string   `json:"target"`
	Ready      bool     `json:"ready"`
	Findings   []string `json:"findings"`
}

type lockRecord struct {
	Capability  string            `json:"capability"`
	Target      string            `json:"target"`
	Method      string            `json:"method"`
	Destination string            `json:"destination"`
	Backup      string            `json:"backup,omitempty"`
	CreatedAt   string            `json:"created_at"`
	Files       map[string]string `json:"files"`
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "agentx:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return usage()
	}

	root, err := repoRoot()
	if err != nil {
		return err
	}

	switch args[0] {
	case "detect":
		if len(args) != 2 {
			return usage()
		}
		return detect(root, args[1])
	case "list":
		return listCapabilities(root)
	case "benchmark":
		return runBenchmark(root)
	case "verify":
		if len(args) < 2 {
			return usage()
		}
		target := flagValue(args[2:], "--target")
		if target == "" {
			return errors.New("verify requires --target <target>")
		}
		result := targetReady(root, args[1], target)
		printJSON(result)
		if !result.Ready {
			return errors.New("target is not ready")
		}
		return nil
	case "plan":
		if len(args) < 4 || args[1] != "install" {
			return usage()
		}
		target := flagValue(args[3:], "--target")
		if target == "" {
			return errors.New("plan install requires --target <target>")
		}
		path, err := planInstall(root, args[2], target)
		if path != "" {
			fmt.Println(path)
		}
		return err
	case "export":
		if len(args) < 2 {
			return usage()
		}
		return deliver(root, args[1], flagValue(args[2:], "--target"), flagValue(args[2:], "--dest"), hasFlag(args[2:], "--yes"), "export")
	case "install":
		if len(args) < 2 {
			return usage()
		}
		return deliver(root, args[1], flagValue(args[2:], "--target"), flagValue(args[2:], "--dest"), hasFlag(args[2:], "--yes"), "install")
	case "rollback":
		if len(args) < 2 {
			return usage()
		}
		return rollback(root, args[1], flagValue(args[2:], "--target"), hasFlag(args[2:], "--yes"))
	default:
		return usage()
	}
}

func usage() error {
	return errors.New("usage: agentx detect <target> | list | benchmark | verify <capability> --target <target> | plan install <capability> --target <target> | export/install <capability> --target <target> --dest <path> --yes | rollback <capability> --target <target> --yes")
}

func runBenchmark(root string) error {
	benchRoot, err := os.MkdirTemp("", "agentx-benchmark-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(benchRoot)
	defer cleanupBenchmarkArtifacts(root)

	report := benchReport{}
	scenarios := []func(string) benchResult{
		benchCleanSkillReady,
		benchMissingBriefBlocks,
		benchMissingReviewBlocks,
		benchMissingStatusBlocks,
		benchNeedsUserDecisionBlocks,
		benchMultipleStatusBlocks,
		benchBlockedReviewBlocks,
		benchReviewPlaceholderBlocks,
		benchUnknownResolutionBlocks,
		benchTargetPlaceholderBlocks,
		benchTargetSymlinkBlocks,
		benchCodexMissingSkillBlocks,
		benchRuntimeGapBlocks,
		benchRuntimeManualPasses,
		benchRuntimeTargetIsolation,
		benchSafetyBlocked,
		benchCursorRulePasses,
		benchBaselineDeviationRecorded,
		benchBaselineDeviationMissingBlocks,
		benchInstallRequiresConfirmation,
		benchPlanBlockedArtifact,
		benchInstallPlanManualRequirements,
		benchDangerousDestinationBlocks,
		benchRepoSubdirDestinationBlocks,
		benchInvalidCapabilityIDBlocks,
		benchInvalidTargetIDBlocks,
		benchExportCopiesReadyArtifact,
		benchRollbackRestoresBackup,
		benchMultiTurnIntakeCoverage,
		benchFinalGenerationGateCoverage,
		benchInstallPlanOnlyCoverage,
		benchMetaSkillRoutingCoverage,
		benchTargetProfileCompleteness,
		benchProfileUnknownClean,
		benchSourceIntakeCoverage,
		benchSafetyAuditCoverage,
		benchPlanReadyArtifact,
		benchSkillMetadataQuality,
		benchSkillOperationalSections,
		benchSkillProgressiveDisclosureQuality,
		benchSkillDirectoryMinimal,
		benchWorkbenchDoesNotDuplicate,
		benchTranslatorConversionLossCoverage,
		benchArchitectDecisionPointCoverage,
		benchBenchmarkMethodSourceCoverage,
		benchRandomizedTargetReadyFuzz,
	}

	for _, scenario := range scenarios {
		result := scenario(root)
		report.Scenarios = append(report.Scenarios, result)
		if result.Passed {
			report.Passed++
		} else {
			report.Failed++
		}
	}
	_ = benchRoot
	printJSON(report)
	if report.Failed > 0 {
		return errors.New("benchmark failed")
	}
	return nil
}

func benchCleanSkillReady(root string) benchResult {
	name := "clean-skill-ready"
	capability := "bench-clean"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "ready", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "ready", result, true)
}

func benchMissingBriefBlocks(root string) benchResult {
	name := "missing-brief-blocks"
	capability := "bench-missing-brief"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	if err := os.Remove(filepath.Join(capabilityDir(root, capability), "brief.md")); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchMissingReviewBlocks(root string) benchResult {
	name := "missing-review-blocks"
	capability := "bench-missing-review"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	if err := os.Remove(filepath.Join(capabilityDir(root, capability), "reviews", "safety-review.md")); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchMissingStatusBlocks(root string) benchResult {
	name := "missing-status-blocks"
	capability := "bench-missing-status"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "semantic-review.md"), []byte("# Semantic Review\n"), 0o644); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchNeedsUserDecisionBlocks(root string) benchResult {
	name := "needs-user-decision-blocks"
	capability := "bench-needs-user-decision"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "portability-review.md"), []byte("Status: needs-user-decision\n"), 0o644); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchMultipleStatusBlocks(root string) benchResult {
	name := "multiple-status-blocks"
	capability := "bench-multiple-status"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	text := "Status: passed\nStatus: needs-user-decision\n"
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "semantic-review.md"), []byte(text), 0o644); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchBlockedReviewBlocks(root string) benchResult {
	name := "blocked-review-blocks"
	capability := "bench-blocked-review"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "semantic-review.md"), []byte("Status: blocked\n"), 0o644); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchReviewPlaceholderBlocks(root string) benchResult {
	name := "review-placeholder-blocks"
	capability := "bench-review-placeholder"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	text := "Status: passed\n\nResidual TODO item\n"
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "semantic-review.md"), []byte(text), 0o644); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchUnknownResolutionBlocks(root string) benchResult {
	name := "unknown-resolution-blocks"
	capability := "bench-unknown"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	text := "Status: passed\n\n- Source: target profile\n- Resolution: Unknown\n"
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "unknown-resolution.md"), []byte(text), 0o644); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchTargetPlaceholderBlocks(root string) benchResult {
	name := "target-placeholder-blocks"
	capability := "bench-placeholder"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "Unknown placeholder", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchTargetSymlinkBlocks(root string) benchResult {
	name := "target-symlink-blocks"
	capability := "bench-target-symlink"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	linkPath := filepath.Join(capabilityDir(root, capability), "targets", "codex", "linked-secret")
	if err := os.Symlink(filepath.Join(os.TempDir(), "agentx-nonexistent-secret"), linkPath); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchCodexMissingSkillBlocks(root string) benchResult {
	name := "codex-missing-skill-blocks"
	capability := "bench-codex-missing-skill"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	if err := os.Remove(targetArtifactPath(root, capability, "codex")); err != nil {
		return benchErr(name, "blocked", err)
	}
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "targets", "codex", "bench.mdc"), []byte("Status: passed\n"), 0o644); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchRuntimeGapBlocks(root string) benchResult {
	name := "runtime-gap-blocks"
	capability := "bench-runtime-gap"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "blocked", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchRuntimeManualPasses(root string) benchResult {
	name := "runtime-manual-passes"
	capability := "bench-runtime-manual"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "manual-transcript", "", false, true, false); err != nil {
		return benchErr(name, "ready", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "ready", result, true)
}

func benchRuntimeTargetIsolation(root string) benchResult {
	name := "runtime-target-isolation"
	capability := "bench-runtime-target-isolation"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "blocked", "", false, true, false); err != nil {
		return benchErr(name, "codex blocked and claude-code ready", err)
	}
	if err := writeBenchCapability(root, capability, "claude-code", "passed", "", false, true, false); err != nil {
		return benchErr(name, "codex blocked and claude-code ready", err)
	}
	codex := targetReady(root, capability, "codex")
	claude := targetReady(root, capability, "claude-code")
	passed := !codex.Ready && claude.Ready
	var findings []string
	if codex.Ready {
		findings = append(findings, "codex unexpectedly ready")
	}
	if !claude.Ready {
		findings = append(findings, "claude-code unexpectedly blocked: "+strings.Join(claude.Findings, "; "))
	}
	return benchResult{Name: name, Expected: "codex blocked and claude-code ready", Passed: passed, Findings: findings}
}

func benchSafetyBlocked(root string) benchResult {
	name := "unsafe-safety-review-blocks"
	capability := "bench-safety"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "safety-review.md"), []byte("Status: blocked\n\nDangerous command\n"), 0o644); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchCursorRulePasses(root string) benchResult {
	name := "cursor-rule-package"
	capability := "bench-cursor"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "cursor", "passed", "", false, true, true); err != nil {
		return benchErr(name, "ready", err)
	}
	result := targetReady(root, capability, "cursor")
	return expectReady(name, "ready", result, true)
}

func benchBaselineDeviationRecorded(root string) benchResult {
	name := "baseline-deviation-recorded"
	capability := "bench-baseline-recorded"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", true, true, false); err != nil {
		return benchErr(name, "ready", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "ready", result, true)
}

func benchBaselineDeviationMissingBlocks(root string) benchResult {
	name := "baseline-deviation-missing-blocks"
	capability := "bench-baseline-missing"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", true, false, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "codex")
	return expectReady(name, "blocked", result, false)
}

func benchInstallRequiresConfirmation(root string) benchResult {
	name := "install-without-confirmation-blocks"
	capability := "bench-install-confirm"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	err := deliver(root, capability, "codex", filepath.Join(os.TempDir(), "agentx-bench-install"), false, "install")
	if err == nil {
		return benchResult{Name: name, Expected: "blocked", Passed: false, Error: "install without --yes unexpectedly succeeded"}
	}
	return benchResult{Name: name, Expected: "blocked", Passed: strings.Contains(err.Error(), "--yes")}
}

func benchPlanBlockedArtifact(root string) benchResult {
	name := "plan-blocked-artifact"
	capability := "bench-plan-blocked"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "blocked", "", false, true, false); err != nil {
		return benchErr(name, "blocked plan", err)
	}
	_, err := planInstall(root, capability, "codex")
	plan := readText(filepath.Join(capabilityDir(root, capability), "install", "codex.plan.md"))
	passed := err != nil && strings.Contains(plan, "Status: blocked")
	return benchResult{Name: name, Expected: "blocked plan", Passed: passed, Error: errString(err)}
}

func benchInstallPlanManualRequirements(root string) benchResult {
	name := "install-plan-manual-requirements"
	capability := "bench-plan-manual"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "manual requirements included", err)
	}
	unknown := "Status: passed\n\n## missing-source-license-file\n\n- Source: source frontmatter\n- Impact: installation\n- Resolution: manual\n- Evidence: license file was not provided.\n- Final artifact impact: install or redistribution requires license confirmation.\n"
	if err := os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "unknown-resolution.md"), []byte(unknown), 0o644); err != nil {
		return benchErr(name, "manual requirements included", err)
	}
	path, err := planInstall(root, capability, "codex")
	plan := readText(path)
	passed := err == nil && strings.Contains(plan, "## Manual Requirements") && strings.Contains(plan, "missing-source-license-file")
	return benchResult{Name: name, Expected: "manual requirements included", Passed: passed, Error: errString(err)}
}

func benchDangerousDestinationBlocks(root string) benchResult {
	name := "dangerous-destination-blocks"
	capability := "bench-dangerous-dest"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	err := deliver(root, capability, "codex", root, true, "export")
	if err == nil {
		return benchResult{Name: name, Expected: "blocked", Passed: false, Error: "export to repository root unexpectedly succeeded"}
	}
	return benchResult{Name: name, Expected: "blocked", Passed: strings.Contains(err.Error(), "dangerous destination"), Error: errString(err)}
}

func benchRepoSubdirDestinationBlocks(root string) benchResult {
	name := "repo-subdir-destination-blocks"
	capability := "bench-repo-subdir-dest"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	dest := filepath.Join(root, "helper", "tmp-agentx-export")
	err := deliver(root, capability, "codex", dest, true, "export")
	if err == nil {
		return benchResult{Name: name, Expected: "blocked", Passed: false, Error: "export to repository subdirectory unexpectedly succeeded"}
	}
	return benchResult{Name: name, Expected: "blocked", Passed: strings.Contains(err.Error(), "dangerous destination"), Error: errString(err)}
}

func benchInvalidCapabilityIDBlocks(root string) benchResult {
	name := "invalid-capability-id-blocks"
	result := targetReady(root, "../escape", "codex")
	return expectReady(name, "blocked", result, false)
}

func benchInvalidTargetIDBlocks(root string) benchResult {
	name := "invalid-target-id-blocks"
	capability := "bench-invalid-target"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "blocked", err)
	}
	result := targetReady(root, capability, "../codex")
	return expectReady(name, "blocked", result, false)
}

func benchExportCopiesReadyArtifact(root string) benchResult {
	name := "export-copies-ready-artifact"
	capability := "bench-export-ready"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	dest, err := os.MkdirTemp("", "agentx-export-*")
	if err != nil {
		return benchErr(name, "exported files", err)
	}
	defer os.RemoveAll(dest)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "exported files", err)
	}
	if err := deliver(root, capability, "codex", dest, true, "export"); err != nil {
		return benchErr(name, "exported files", err)
	}
	passed := exists(filepath.Join(dest, "SKILL.md")) && exists(filepath.Join(capabilityDir(root, capability), "lock.json"))
	return benchResult{Name: name, Expected: "exported files", Passed: passed}
}

func benchRollbackRestoresBackup(root string) benchResult {
	name := "rollback-restores-backup"
	capability := "bench-rollback"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	dest, err := os.MkdirTemp("", "agentx-rollback-*")
	if err != nil {
		return benchErr(name, "restored backup", err)
	}
	defer os.RemoveAll(dest)
	original := filepath.Join(dest, "existing.txt")
	if err := os.WriteFile(original, []byte("original\n"), 0o644); err != nil {
		return benchErr(name, "restored backup", err)
	}
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "restored backup", err)
	}
	if err := deliver(root, capability, "codex", dest, true, "install"); err != nil {
		return benchErr(name, "restored backup", err)
	}
	if err := rollback(root, capability, "codex", true); err != nil {
		return benchErr(name, "restored backup", err)
	}
	restored := readText(original)
	return benchResult{Name: name, Expected: "restored backup", Passed: restored == "original\n"}
}

func benchMultiTurnIntakeCoverage(root string) benchResult {
	name := "multi-turn-intake-coverage"
	spec := readText(filepath.Join(root, "spec", "01-ai-native-workflow.md"))
	workbench := readText(filepath.Join(root, "skills", "agentx-workbench", "SKILL.md"))
	findings := requireContains(spec+workbench, []string{
		"Do you want to add more input?",
		"\"no\", \"done\", \"不\", or \"不需要\"",
		".agentx/output/capabilities/<id>/",
		"intake.md",
		"open-questions.md",
	})
	return benchFromFindings(name, "covered", findings)
}

func benchFinalGenerationGateCoverage(root string) benchResult {
	name := "final-generation-gate-coverage"
	spec := readText(filepath.Join(root, "spec", "01-ai-native-workflow.md"))
	findings := requireContains(spec, []string{
		"The user closes intake and confirms generation.",
		"The user asks it to generate an early draft now.",
	})
	if strings.Contains(spec, "has "+"enough information") {
		findings = append(findings, "final generation can bypass intake close")
	}
	return benchFromFindings(name, "covered", findings)
}

func benchInstallPlanOnlyCoverage(root string) benchResult {
	name := "install-plan-only-coverage"
	text := readText(filepath.Join(root, "spec", "08-installation-delivery.md")) + readText(filepath.Join(root, "skills", "agentx-install-planner", "SKILL.md"))
	findings := requireContains(text, []string{
		"The default delivery mode is planning only.",
		"explicitly asks it to proceed",
		"Default mode is plan-only.",
		"requires explicit --yes confirmation",
	})
	return benchFromFindings(name, "covered", findings)
}

func benchMetaSkillRoutingCoverage(root string) benchResult {
	name := "meta-skill-routing-coverage"
	workbench := readText(filepath.Join(root, "skills", "agentx-workbench", "SKILL.md"))
	findings := requireContains(workbench, []string{
		"agentx-capability-architect",
		"agentx-capability-translator",
		"agentx-capability-reviewer",
		"agentx-safety-auditor",
		"agentx-benchmark-designer",
		"agentx-install-planner",
	})
	return benchFromFindings(name, "covered", findings)
}

func benchTargetProfileCompleteness(root string) benchResult {
	name := "target-profile-completeness"
	var findings []string
	targetRoot := filepath.Join(root, "targets")
	entries, err := os.ReadDir(targetRoot)
	if err != nil {
		return benchErr(name, "complete", err)
	}
	required := []string{"profile.md", "creator-baseline.md", "package-layout.md", "install.md", "review-checklist.md"}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		for _, file := range required {
			path := filepath.Join(targetRoot, entry.Name(), file)
			if !exists(path) {
				findings = append(findings, "missing "+filepath.ToSlash(filepath.Join("targets", entry.Name(), file)))
			}
		}
	}
	return benchFromFindings(name, "complete", findings)
}

func benchProfileUnknownClean(root string) benchResult {
	name := "profile-unknown-clean"
	var findings []string
	for _, dir := range []string{"targets", "models"} {
		_ = filepath.WalkDir(filepath.Join(root, dir), func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
				return nil
			}
			text := readText(path)
			for _, line := range strings.Split(text, "\n") {
				trimmed := strings.TrimSpace(line)
				if strings.HasPrefix(trimmed, "Status: Unknown") || strings.HasSuffix(trimmed, ": Unknown") || trimmed == "Unknown" {
					rel, _ := filepath.Rel(root, path)
					findings = append(findings, "unresolved profile Unknown in "+filepath.ToSlash(rel))
				}
			}
			return nil
		})
	}
	return benchFromFindings(name, "clean", findings)
}

func benchSourceIntakeCoverage(root string) benchResult {
	name := "source-intake-coverage"
	text := readText(filepath.Join(root, "spec", "06-translation-workflow.md")) + readText(filepath.Join(root, "skills", "agentx-capability-translator", "SKILL.md"))
	findings := requireContains(text, []string{
		"path",
		"URL",
		"repository",
		"directory",
		"pasted content",
		"do not ask the user to paste locally readable files",
	})
	return benchFromFindings(name, "covered", findings)
}

func benchSafetyAuditCoverage(root string) benchResult {
	name := "safety-audit-coverage"
	text := readText(filepath.Join(root, "spec", "07-review-safety-benchmark.md")) + readText(filepath.Join(root, "skills", "agentx-safety-auditor", "SKILL.md"))
	findings := requireContains(text, []string{
		"Prompt injection",
		"Secret",
		"Credential",
		"Dangerous shell",
		"Network exfiltration",
		"Obfuscated scripts",
		"Excessive",
	})
	return benchFromFindings(name, "covered", findings)
}

func benchPlanReadyArtifact(root string) benchResult {
	name := "plan-ready-artifact"
	capability := "bench-plan-ready"
	cleanupCapability(root, capability)
	defer cleanupCapability(root, capability)
	if err := writeBenchCapability(root, capability, "codex", "passed", "", false, true, false); err != nil {
		return benchErr(name, "ready plan", err)
	}
	path, err := planInstall(root, capability, "codex")
	plan := readText(path)
	passed := err == nil && strings.Contains(plan, "Status: ready") && strings.Contains(plan, "Default")
	if !passed && err == nil {
		passed = strings.Contains(plan, "Status: ready") && strings.Contains(plan, "User Confirmation")
	}
	return benchResult{Name: name, Expected: "ready plan", Passed: passed, Error: errString(err)}
}

func benchSkillMetadataQuality(root string) benchResult {
	name := "skill-metadata-quality"
	var findings []string
	for _, path := range skillFiles(root) {
		text := readText(path)
		meta := parseFrontmatter(text)
		rel, _ := filepath.Rel(root, path)
		if meta["name"] == "" {
			findings = append(findings, filepath.ToSlash(rel)+" missing frontmatter name")
		}
		if len(meta["description"]) < 60 {
			findings = append(findings, filepath.ToSlash(rel)+" description is too short or missing")
		}
		if !containsAny(strings.ToLower(meta["description"]), []string{"create", "translate", "review", "audit", "design", "produce", "route"}) {
			findings = append(findings, filepath.ToSlash(rel)+" description lacks an action verb")
		}
	}
	return benchFromFindings(name, "high-signal skill metadata", findings)
}

func benchSkillOperationalSections(root string) benchResult {
	name := "skill-operational-sections"
	var findings []string
	markers := []string{"## Procedure", "## Routing Procedure", "Before producing", "When reviewing", "## Required Files"}
	for _, path := range skillFiles(root) {
		text := readText(path)
		rel, _ := filepath.Rel(root, path)
		if !containsAny(text, markers) {
			findings = append(findings, filepath.ToSlash(rel)+" lacks an operational procedure section")
		}
	}
	return benchFromFindings(name, "actionable skill procedures", findings)
}

func benchSkillProgressiveDisclosureQuality(root string) benchResult {
	name := "skill-progressive-disclosure-quality"
	required := map[string][]string{
		"agentx-workbench":             {"Do not duplicate the full instructions", "Workflow Map"},
		"agentx-capability-architect":  {"Follow `spec/01-ai-native-workflow.md`", "Follow `spec/05-capability-brief.md`"},
		"agentx-capability-translator": {"Read every requested `targets/<target-id>/` profile", "official creator baseline"},
		"agentx-capability-reviewer":   {"reviews/unknown-resolution.md", "Final target files do not contain"},
		"agentx-safety-auditor":        {"prompt injection", "network exfiltration"},
		"agentx-benchmark-designer":    {"runtime-benchmark.<target-id>.md", "Final delivery requires"},
		"agentx-install-planner":       {"target-ready gate", "Default mode is plan-only"},
	}
	var findings []string
	for skill, needles := range required {
		path := filepath.Join(root, "skills", skill, "SKILL.md")
		if !exists(path) {
			findings = append(findings, "missing skill "+skill)
			continue
		}
		for _, finding := range requireContains(readText(path), needles) {
			findings = append(findings, skill+": "+finding)
		}
	}
	return benchFromFindings(name, "progressive disclosure enforced", findings)
}

func benchSkillDirectoryMinimal(root string) benchResult {
	name := "skill-directory-minimal"
	var findings []string
	entries, err := os.ReadDir(filepath.Join(root, "skills"))
	if err != nil {
		return benchErr(name, "minimal skill directories", err)
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		dir := filepath.Join(root, "skills", entry.Name())
		_ = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			if filepath.Base(path) != "SKILL.md" {
				rel, _ := filepath.Rel(root, path)
				findings = append(findings, filepath.ToSlash(rel)+" is extra skill payload")
			}
			return nil
		})
	}
	return benchFromFindings(name, "minimal skill payloads", findings)
}

func benchWorkbenchDoesNotDuplicate(root string) benchResult {
	name := "workbench-does-not-duplicate"
	text := readText(filepath.Join(root, "skills", "agentx-workbench", "SKILL.md"))
	var findings []string
	if lineCount(text) > 90 {
		findings = append(findings, "workbench is too large for a routing skill")
	}
	for _, forbidden := range []string{"Preserved, Adapted, Degraded", "Prompt injection.", "Status: manual-transcript"} {
		if strings.Contains(text, forbidden) {
			findings = append(findings, "workbench duplicates specialized instruction: "+forbidden)
		}
	}
	return benchFromFindings(name, "routing-only workbench", findings)
}

func benchTranslatorConversionLossCoverage(root string) benchResult {
	name := "translator-conversion-loss-coverage"
	text := readText(filepath.Join(root, "skills", "agentx-capability-translator", "SKILL.md")) + readText(filepath.Join(root, "spec", "06-translation-workflow.md"))
	findings := requireContains(text, []string{
		"conversion loss report",
		"Preserved",
		"Adapted",
		"Degraded",
		"Dropped",
		"Manual Setup Required",
		"Risks",
	})
	return benchFromFindings(name, "conversion loss coverage", findings)
}

func benchArchitectDecisionPointCoverage(root string) benchResult {
	name := "architect-decision-point-coverage"
	text := readText(filepath.Join(root, "skills", "agentx-capability-architect", "SKILL.md")) + readText(filepath.Join(root, "spec", "01-ai-native-workflow.md"))
	findings := requireContains(text, []string{
		"After each intake turn",
		"After intake closes",
		"decision points",
		"open-questions.md",
		"early draft",
	})
	return benchFromFindings(name, "decision points before generation", findings)
}

func benchBenchmarkMethodSourceCoverage(root string) benchResult {
	name := "benchmark-method-source-coverage"
	text := readText(filepath.Join(root, "references", "standards", "benchmark-method.md"))
	findings := requireContains(text, []string{
		"https://www.skills.sh/docs",
		"https://www.skill-issue.sh/",
		"https://www.oasb.ai/benchmark",
		"https://agentdojo.spylab.ai/",
		"https://www.skillsbench.ai/",
		"positive cases",
		"negative cases",
		"decoys",
		"routing confusion tests",
	})
	return benchFromFindings(name, "benchmark method grounded in references", findings)
}

func benchRandomizedTargetReadyFuzz(root string) benchResult {
	name := "randomized-target-ready-fuzz"
	seed := int64(20260527)
	if value := os.Getenv("AGENTX_BENCH_SEED"); value != "" {
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return benchErr(name, "valid random seed", err)
		}
		seed = parsed
	}
	rnd := rand.New(rand.NewSource(seed))
	targets := []string{"codex", "cursor"}
	mutations := []struct {
		name  string
		apply func(string, string) error
	}{
		{
			name: "missing-brief",
			apply: func(capability, target string) error {
				return os.Remove(filepath.Join(capabilityDir(root, capability), "brief.md"))
			},
		},
		{
			name: "missing-review",
			apply: func(capability, target string) error {
				return os.Remove(filepath.Join(capabilityDir(root, capability), "reviews", "semantic-review.md"))
			},
		},
		{
			name: "blank-review-status",
			apply: func(capability, target string) error {
				return os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "safety-review.md"), []byte("# Safety Review\n"), 0o644)
			},
		},
		{
			name: "needs-user-decision",
			apply: func(capability, target string) error {
				return os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "portability-review.md"), []byte("Status: needs-user-decision\n"), 0o644)
			},
		},
		{
			name: "runtime-blocked",
			apply: func(capability, target string) error {
				return os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", runtimeBenchmarkName(target)), []byte("Status: blocked\n"), 0o644)
			},
		},
		{
			name: "unknown-tbd",
			apply: func(capability, target string) error {
				text := "Status: passed\n\n- Source: target profile\n- Resolution: TBD\n"
				return os.WriteFile(filepath.Join(capabilityDir(root, capability), "reviews", "unknown-resolution.md"), []byte(text), 0o644)
			},
		},
		{
			name: "placeholder-in-target",
			apply: func(capability, target string) error {
				return os.WriteFile(targetArtifactPath(root, capability, target), []byte("Status: passed\n\nUnknown target behavior\n"), 0o644)
			},
		},
		{
			name: "baseline-deviation-without-record",
			apply: func(capability, target string) error {
				reviews := filepath.Join(capabilityDir(root, capability), "reviews")
				if err := os.WriteFile(filepath.Join(reviews, "semantic-review.md"), []byte("Status: passed\nBaseline deviation: yes\n"), 0o644); err != nil {
					return err
				}
				return os.Remove(filepath.Join(reviews, "baseline-deviations.md"))
			},
		},
	}

	var findings []string
	for i := 0; i < 40; i++ {
		target := targets[rnd.Intn(len(targets))]
		capability := fmt.Sprintf("bench-fuzz-%02d", i)
		cleanupCapability(root, capability)
		cursor := target == "cursor"
		if err := writeBenchCapability(root, capability, target, "passed", "", false, true, cursor); err != nil {
			findings = append(findings, fmt.Sprintf("%s setup failed: %v", capability, err))
			continue
		}

		mutationCount := rnd.Intn(4)
		var labels []string
		for j := 0; j < mutationCount; j++ {
			mutation := mutations[rnd.Intn(len(mutations))]
			labels = append(labels, mutation.name)
			if err := mutation.apply(capability, target); err != nil && !os.IsNotExist(err) {
				findings = append(findings, fmt.Sprintf("%s mutation %s failed: %v", capability, mutation.name, err))
			}
		}

		result := targetReady(root, capability, target)
		wantReady := mutationCount == 0
		if result.Ready != wantReady {
			findings = append(findings, fmt.Sprintf("%s target=%s mutations=%v ready=%v want=%v findings=%v", capability, target, labels, result.Ready, wantReady, result.Findings))
		}
		cleanupCapability(root, capability)
	}
	return benchFromFindings(name, "fixed-seed randomized target-ready edge cases", findings)
}

func targetArtifactPath(root, capability, target string) string {
	if target == "cursor" {
		return filepath.Join(capabilityDir(root, capability), "targets", target, ".cursor", "rules", "bench.mdc")
	}
	return filepath.Join(capabilityDir(root, capability), "targets", target, "SKILL.md")
}

func runtimeBenchmarkName(target string) string {
	return "runtime-benchmark." + target + ".md"
}

func writeBenchCapability(root, capability, target, runtimeStatus, targetExtra string, baselineDeviation, writeBaselineFile, cursor bool) error {
	base := capabilityDir(root, capability)
	reviews := filepath.Join(base, "reviews")
	targetDir := filepath.Join(base, "targets", target)
	if cursor {
		targetDir = filepath.Join(targetDir, ".cursor", "rules")
	}
	if err := os.MkdirAll(base, 0o755); err != nil {
		return err
	}
	for name, text := range map[string]string{
		"intake.md":         "# Intake\n\n- Benchmark requirement source.\n",
		"open-questions.md": "# Open Questions\n\nStatus: resolved\n",
		"brief.md":          "# Capability Brief\n\nStatus: ready\n",
	} {
		if err := os.WriteFile(filepath.Join(base, name), []byte(text), 0o644); err != nil {
			return err
		}
	}
	if err := os.MkdirAll(reviews, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}
	deviation := "no"
	if baselineDeviation {
		deviation = "yes"
	}
	reviewText := "Status: passed\nBaseline deviation: " + deviation + "\n"
	for _, name := range []string{"semantic-review.md", "portability-review.md"} {
		if err := os.WriteFile(filepath.Join(reviews, name), []byte(reviewText), 0o644); err != nil {
			return err
		}
	}
	for _, name := range []string{"safety-review.md", "benchmark-plan.md"} {
		if err := os.WriteFile(filepath.Join(reviews, name), []byte("Status: passed\n"), 0o644); err != nil {
			return err
		}
	}
	if err := os.WriteFile(filepath.Join(reviews, runtimeBenchmarkName(target)), []byte("Status: "+runtimeStatus+"\n"), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(reviews, "unknown-resolution.md"), []byte("Status: passed\n\n- Source: target profile\n- Resolution: verified\n"), 0o644); err != nil {
		return err
	}
	if writeBaselineFile {
		if err := os.WriteFile(filepath.Join(reviews, "baseline-deviations.md"), []byte("Status: passed\n"), 0o644); err != nil {
			return err
		}
	}
	targetText := "---\nname: bench\n" + "description: Use for AgentX benchmark.\n---\n\n# Benchmark\n" + targetExtra + "\n"
	targetPath := filepath.Join(targetDir, "SKILL.md")
	if cursor {
		targetText = "---\ndescription: Use for AgentX benchmark.\nalwaysApply: false\n---\n\n# Benchmark Cursor Rule\n" + targetExtra + "\n"
		targetPath = filepath.Join(targetDir, "bench.mdc")
	}
	return os.WriteFile(targetPath, []byte(targetText), 0o644)
}

func cleanupCapability(root, capability string) {
	_ = os.RemoveAll(capabilityDir(root, capability))
}

func cleanupBenchmarkArtifacts(root string) {
	base := filepath.Join(root, ".agentx", "output", "capabilities")
	entries, err := os.ReadDir(base)
	if err == nil {
		for _, entry := range entries {
			if entry.IsDir() && strings.HasPrefix(entry.Name(), "bench-") {
				_ = os.RemoveAll(filepath.Join(base, entry.Name()))
			}
		}
	}
	backupBase := filepath.Join(root, ".agentx", "backups")
	backupEntries, err := os.ReadDir(backupBase)
	if err == nil {
		for _, entry := range backupEntries {
			if entry.IsDir() && strings.HasPrefix(entry.Name(), "bench-") {
				_ = os.RemoveAll(filepath.Join(backupBase, entry.Name()))
			}
		}
	}
	removeIfEmpty(base)
	removeIfEmpty(filepath.Dir(base))
	removeIfEmpty(filepath.Join(root, ".agentx", "output"))
	removeIfEmpty(backupBase)
	removeIfEmpty(filepath.Join(root, ".agentx"))
}

func removeIfEmpty(path string) {
	entries, err := os.ReadDir(path)
	if err == nil && len(entries) == 0 {
		_ = os.Remove(path)
	}
}

func expectReady(name, expected string, result gateResult, want bool) benchResult {
	return benchResult{
		Name:     name,
		Expected: expected,
		Passed:   result.Ready == want,
		Findings: result.Findings,
	}
}

func benchErr(name, expected string, err error) benchResult {
	return benchResult{Name: name, Expected: expected, Passed: false, Error: err.Error()}
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func requireContains(text string, needles []string) []string {
	var findings []string
	for _, needle := range needles {
		if !strings.Contains(text, needle) {
			findings = append(findings, "missing required text: "+needle)
		}
	}
	return findings
}

func benchFromFindings(name, expected string, findings []string) benchResult {
	return benchResult{
		Name:     name,
		Expected: expected,
		Passed:   len(findings) == 0,
		Findings: findings,
	}
}

func skillFiles(root string) []string {
	var files []string
	entries, err := os.ReadDir(filepath.Join(root, "skills"))
	if err != nil {
		return files
	}
	for _, entry := range entries {
		if entry.IsDir() {
			path := filepath.Join(root, "skills", entry.Name(), "SKILL.md")
			if exists(path) {
				files = append(files, path)
			}
		}
	}
	sort.Strings(files)
	return files
}

func parseFrontmatter(text string) map[string]string {
	meta := map[string]string{}
	if !strings.HasPrefix(text, "---\n") {
		return meta
	}
	rest := strings.TrimPrefix(text, "---\n")
	end := strings.Index(rest, "\n---")
	if end < 0 {
		return meta
	}
	for _, line := range strings.Split(rest[:end], "\n") {
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		meta[strings.TrimSpace(key)] = strings.Trim(strings.TrimSpace(value), `"'`)
	}
	return meta
}

func containsAny(text string, needles []string) bool {
	for _, needle := range needles {
		if strings.Contains(text, needle) {
			return true
		}
	}
	return false
}

func lineCount(text string) int {
	if text == "" {
		return 0
	}
	return strings.Count(text, "\n") + 1
}

func repoRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if exists(filepath.Join(wd, "spec")) && exists(filepath.Join(wd, "skills")) && exists(filepath.Join(wd, "targets")) {
			return wd, nil
		}
		next := filepath.Dir(wd)
		if next == wd {
			return "", errors.New("could not locate AgentX repository root")
		}
		wd = next
	}
}

func detect(root, target string) error {
	if err := validateTargetID(root, target); err != nil {
		return err
	}
	info := map[string]any{
		"target":       target,
		"profile_path": filepath.Join(root, "targets", target),
		"profile_ok":   exists(filepath.Join(root, "targets", target, "profile.md")),
	}
	switch target {
	case "codex":
		home := os.Getenv("CODEX_HOME")
		if home == "" {
			home = filepath.Join(userHome(), ".codex")
		}
		info["default_skill_path"] = filepath.Join(home, "skills")
	case "claude-code":
		info["project_skill_path"] = ".claude/skills"
		info["user_skill_path"] = filepath.Join(userHome(), ".claude", "skills")
	case "copilot":
		info["repository_skill_path"] = ".github/skills"
		info["user_skill_path"] = filepath.Join(userHome(), ".copilot", "skills")
	case "cursor":
		info["project_rule_path"] = ".cursor/rules"
	case "hermes":
		info["user_skill_path"] = filepath.Join(userHome(), ".hermes", "skills")
	case "openclaw":
		info["installer"] = "openclaw skill install"
	}
	printJSON(info)
	return nil
}

func listCapabilities(root string) error {
	base := filepath.Join(root, ".agentx", "output", "capabilities")
	entries, err := os.ReadDir(base)
	if os.IsNotExist(err) {
		fmt.Println("[]")
		return nil
	}
	if err != nil {
		return err
	}
	var caps []string
	for _, entry := range entries {
		if entry.IsDir() {
			caps = append(caps, entry.Name())
		}
	}
	sort.Strings(caps)
	printJSON(caps)
	return nil
}

func targetReady(root, capability, target string) gateResult {
	result := gateResult{Capability: capability, Target: target, Ready: true}
	if err := validateCapabilityID(capability); err != nil {
		result.add(err.Error())
		return result
	}
	if err := validateTargetID(root, target); err != nil {
		result.add(err.Error())
		return result
	}
	base := capabilityDir(root, capability)
	targetDir := filepath.Join(base, "targets", target)
	reviews := filepath.Join(base, "reviews")

	for _, name := range []string{"intake.md", "open-questions.md", "brief.md"} {
		if !exists(filepath.Join(base, name)) {
			result.add("missing workflow artifact: " + name)
		}
	}

	required := []string{
		"semantic-review.md",
		"portability-review.md",
		"safety-review.md",
		"benchmark-plan.md",
		"unknown-resolution.md",
	}
	for _, name := range required {
		path := filepath.Join(reviews, name)
		if !exists(path) {
			result.add("missing review: " + name)
			continue
		}
		if isSymlink(path) {
			result.add("review file must not be a symlink: " + name)
			continue
		}
		text := readText(path)
		statuses := statusValues(text)
		if len(statuses) != 1 {
			result.add(name + " must contain exactly one Status line")
		}
		if hasBlockedStatus(text) {
			result.add("blocked review: " + name)
		}
		if len(statuses) != 1 || statuses[0] != "passed" {
			result.add(name + " must be Status: passed")
		}
		if name != "unknown-resolution.md" {
			if err := scanPlaceholders(path); err != nil {
				result.add(err.Error())
			}
		}
	}

	runtimeName := runtimeBenchmarkName(target)
	runtimePath := filepath.Join(reviews, runtimeName)
	if !exists(runtimePath) {
		result.add("missing review: " + runtimeName)
	} else if isSymlink(runtimePath) {
		result.add("review file must not be a symlink: " + runtimeName)
	} else {
		text := readText(runtimePath)
		statuses := statusValues(text)
		if len(statuses) != 1 {
			result.add(runtimeName + " must contain exactly one Status line")
		}
		if hasBlockedStatus(text) {
			result.add("blocked review: " + runtimeName)
		}
		if len(statuses) != 1 || (statuses[0] != "passed" && statuses[0] != "manual-transcript") {
			result.add(runtimeName + " must be Status: passed or Status: manual-transcript")
		}
		if err := scanPlaceholders(runtimePath); err != nil {
			result.add(err.Error())
		}
	}

	if baselineDeviationRequired(reviews) {
		path := filepath.Join(reviews, "baseline-deviations.md")
		if !exists(path) {
			result.add("missing review: baseline-deviations.md")
		} else {
			text := readText(path)
			if hasBlockedStatus(text) {
				result.add("blocked review: baseline-deviations.md")
			}
			statuses := statusValues(text)
			if len(statuses) != 1 {
				result.add("baseline-deviations.md must contain exactly one Status line")
			}
			if len(statuses) != 1 || statuses[0] != "passed" {
				result.add("baseline-deviations.md must be Status: passed")
			}
			if err := scanPlaceholders(path); err != nil {
				result.add(err.Error())
			}
		}
	}

	unknownResolution := readText(filepath.Join(reviews, "unknown-resolution.md"))
	if unknownResolution != "" {
		if hasBlockedStatus(unknownResolution) {
			result.add("unknown-resolution.md is blocked")
		}
		for _, line := range strings.Split(unknownResolution, "\n") {
			trimmed := strings.TrimSpace(strings.ToLower(line))
			if strings.HasPrefix(trimmed, "- resolution:") || strings.HasPrefix(trimmed, "resolution:") {
				if strings.Contains(trimmed, "unknown") || strings.Contains(trimmed, "tbd") || strings.Contains(trimmed, "todo") || strings.HasSuffix(trimmed, ":") {
					result.add("unknown-resolution.md has unresolved resolution: " + strings.TrimSpace(line))
				}
			}
		}
	}

	if !exists(targetDir) {
		result.add("missing target artifact directory: " + filepath.ToSlash(filepath.Join(".agentx/output/capabilities", capability, "targets", target)))
	} else {
		if err := scanSymlinks(targetDir); err != nil {
			result.add(err.Error())
		}
		if err := validateTargetShape(targetDir, target); err != nil {
			result.add(err.Error())
		}
		if err := scanPlaceholders(targetDir); err != nil {
			result.add(err.Error())
		}
	}

	result.Ready = len(result.Findings) == 0
	return result
}

func (r *gateResult) add(finding string) {
	r.Findings = append(r.Findings, finding)
	r.Ready = false
}

func planInstall(root, capability, target string) (string, error) {
	if err := validateCapabilityID(capability); err != nil {
		return "", err
	}
	if err := validateTargetID(root, target); err != nil {
		return "", err
	}
	result := targetReady(root, capability, target)
	planDir := filepath.Join(capabilityDir(root, capability), "install")
	if err := os.MkdirAll(planDir, 0o755); err != nil {
		return "", err
	}
	status := "ready"
	if !result.Ready {
		status = "blocked"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Install Plan: %s -> %s\n\n", capability, target)
	fmt.Fprintf(&b, "Status: %s\n\n", status)
	if !result.Ready {
		b.WriteString("## Blocked Evidence\n\n")
		for _, finding := range result.Findings {
			fmt.Fprintf(&b, "- %s\n", finding)
		}
		b.WriteString("\n")
	}
	if entries := manualResolutionEntries(filepath.Join(capabilityDir(root, capability), "reviews", "unknown-resolution.md")); len(entries) > 0 {
		b.WriteString("## Manual Requirements\n\n")
		for _, entry := range entries {
			fmt.Fprintf(&b, "- %s\n", entry)
		}
		b.WriteString("\n")
	}
	b.WriteString("## Source Artifact\n\n")
	fmt.Fprintf(&b, "`.agentx/output/capabilities/%s/targets/%s/`\n\n", capability, target)
	b.WriteString("## Target Runtime\n\n")
	fmt.Fprintf(&b, "%s\n\n", target)
	b.WriteString("## Method\n\nplan-only\n\n## Destination\n\nRequires user confirmation.\n\n## Changes\n\nNo filesystem changes in plan mode.\n\n## Backup\n\nRequired before install.\n\n## Verification\n\nRun `agentx verify` after delivery.\n\n## Rollback\n\nUse backup created during install.\n\n## User Confirmation\n\nRequired before install or export.\n")
	path := filepath.Join(planDir, target+".plan.md")
	if err := os.WriteFile(path, []byte(b.String()), 0o644); err != nil {
		return "", err
	}
	if !result.Ready {
		return path, errors.New("target is blocked; wrote blocked plan")
	}
	return path, nil
}

func deliver(root, capability, target, dest string, yes bool, method string) error {
	if err := validateCapabilityID(capability); err != nil {
		return err
	}
	if target == "" {
		return errors.New(method + " requires --target <target>")
	}
	if err := validateTargetID(root, target); err != nil {
		return err
	}
	if dest == "" {
		return errors.New(method + " requires --dest <path>")
	}
	if !yes {
		return errors.New(method + " requires explicit --yes confirmation")
	}
	src := filepath.Join(capabilityDir(root, capability), "targets", target)
	if err := validateDestination(root, src, dest); err != nil {
		return err
	}
	result := targetReady(root, capability, target)
	if !result.Ready {
		printJSON(result)
		return errors.New("target is not ready")
	}
	backupRoot := filepath.Join(root, ".agentx", "backups", capability, target, time.Now().UTC().Format("20060102T150405Z"))
	backupPath := ""
	if exists(dest) {
		if err := copyTree(dest, backupRoot); err != nil {
			return err
		}
		backupPath = backupRoot
	}
	if method == "export" || method == "install" {
		if err := os.RemoveAll(dest); err != nil {
			return err
		}
		if err := copyTree(src, dest); err != nil {
			return err
		}
	}
	files, err := hashTree(dest)
	if err != nil {
		return err
	}
	record := lockRecord{
		Capability:  capability,
		Target:      target,
		Method:      method,
		Destination: dest,
		Backup:      backupPath,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
		Files:       files,
	}
	data, _ := json.MarshalIndent(record, "", "  ")
	return os.WriteFile(filepath.Join(capabilityDir(root, capability), "lock.json"), append(data, '\n'), 0o644)
}

func rollback(root, capability, target string, yes bool) error {
	if err := validateCapabilityID(capability); err != nil {
		return err
	}
	if target == "" {
		return errors.New("rollback requires --target <target>")
	}
	if err := validateTargetID(root, target); err != nil {
		return err
	}
	if !yes {
		return errors.New("rollback requires explicit --yes confirmation")
	}
	lockPath := filepath.Join(capabilityDir(root, capability), "lock.json")
	data, err := os.ReadFile(lockPath)
	if err != nil {
		return err
	}
	var record lockRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return err
	}
	if record.Target != target {
		return errors.New("lock target does not match rollback target")
	}
	if record.Destination == "" {
		return errors.New("lock file has no destination")
	}
	if record.Backup == "" {
		return errors.New("no backups found")
	}
	if err := validateDestination(root, filepath.Join(capabilityDir(root, capability), "targets", target), record.Destination); err != nil {
		return err
	}
	if !exists(record.Backup) {
		return errors.New("backup path does not exist")
	}
	if err := os.RemoveAll(record.Destination); err != nil {
		return err
	}
	if err := copyTree(record.Backup, record.Destination); err != nil {
		return err
	}
	return nil
}

func capabilityDir(root, capability string) string {
	return filepath.Join(root, ".agentx", "output", "capabilities", capability)
}

func validateCapabilityID(capability string) error {
	if capability == "" {
		return errors.New("invalid capability id: empty")
	}
	if filepath.IsAbs(capability) || strings.Contains(capability, "/") || strings.Contains(capability, string(filepath.Separator)) {
		return errors.New("invalid capability id: paths are not allowed")
	}
	if capability == "." || capability == ".." || strings.Contains(capability, "..") {
		return errors.New("invalid capability id: parent traversal is not allowed")
	}
	for _, r := range capability {
		if !(r == '-' || r == '_' || r == '.' || r >= '0' && r <= '9' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			return errors.New("invalid capability id: only letters, numbers, dot, dash, and underscore are allowed")
		}
	}
	return nil
}

func validateTargetID(root, target string) error {
	if target == "" {
		return errors.New("invalid target id: empty")
	}
	if filepath.IsAbs(target) || strings.Contains(target, "/") || strings.Contains(target, string(filepath.Separator)) {
		return errors.New("invalid target id: paths are not allowed")
	}
	if target == "." || target == ".." || strings.Contains(target, "..") {
		return errors.New("invalid target id: parent traversal is not allowed")
	}
	if !exists(filepath.Join(root, "targets", target, "profile.md")) {
		return errors.New("invalid target id: target profile not found")
	}
	return nil
}

func hasBlockedStatus(text string) bool {
	return hasStatus(text, "blocked")
}

func hasStatus(text, want string) bool {
	want = strings.ToLower(want)
	for _, status := range statusValues(text) {
		if status == want {
			return true
		}
	}
	return false
}

func statusValues(text string) []string {
	var values []string
	for _, line := range strings.Split(text, "\n") {
		trimmed := strings.TrimSpace(strings.ToLower(line))
		if value, ok := strings.CutPrefix(trimmed, "status:"); ok {
			values = append(values, strings.TrimSpace(value))
		}
	}
	return values
}

func baselineDeviationRequired(reviews string) bool {
	for _, name := range []string{"semantic-review.md", "portability-review.md", "safety-review.md"} {
		text := strings.ToLower(readText(filepath.Join(reviews, name)))
		for _, line := range strings.Split(text, "\n") {
			if strings.TrimSpace(line) == "baseline deviation: yes" {
				return true
			}
		}
	}
	return false
}

func manualResolutionEntries(path string) []string {
	text := readText(path)
	if text == "" {
		return nil
	}
	var entries []string
	current := ""
	var block []string
	flush := func() {
		if current == "" || len(block) == 0 {
			block = nil
			return
		}
		joined := strings.ToLower(strings.Join(block, "\n"))
		if strings.Contains(joined, "resolution: manual") {
			summary := current
			for _, line := range block {
				trimmed := strings.TrimSpace(line)
				lower := strings.ToLower(trimmed)
				if strings.HasPrefix(lower, "- final artifact impact:") {
					_, value, _ := strings.Cut(trimmed, ":")
					summary += " - " + strings.TrimSpace(value)
					break
				}
			}
			entries = append(entries, summary)
		}
		block = nil
	}
	for _, line := range strings.Split(text, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "## ") {
			flush()
			current = strings.TrimSpace(strings.TrimLeft(trimmed, "# "))
			continue
		}
		if current != "" {
			block = append(block, line)
		}
	}
	flush()
	return entries
}

func scanPlaceholders(root string) error {
	var findings []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		text := readText(path)
		for _, token := range []string{"Unknown", "TBD", "TODO"} {
			if strings.Contains(text, token) {
				rel, _ := filepath.Rel(root, path)
				findings = append(findings, fmt.Sprintf("unresolved placeholder %q in %s", token, filepath.ToSlash(rel)))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(findings) > 0 {
		return errors.New(strings.Join(findings, "; "))
	}
	return nil
}

func scanSymlinks(root string) error {
	var findings []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Type()&os.ModeSymlink != 0 {
			rel, _ := filepath.Rel(root, path)
			findings = append(findings, "symlink is not allowed in target artifact: "+filepath.ToSlash(rel))
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(findings) > 0 {
		return errors.New(strings.Join(findings, "; "))
	}
	return nil
}

func isSymlink(path string) bool {
	info, err := os.Lstat(path)
	return err == nil && info.Mode()&os.ModeSymlink != 0
}

func validateTargetShape(targetDir, target string) error {
	switch target {
	case "codex", "claude-code", "openclaw", "hermes":
		if !exists(filepath.Join(targetDir, "SKILL.md")) {
			return errors.New(target + " target artifact must include SKILL.md")
		}
	case "cursor":
		rulesDir := filepath.Join(targetDir, ".cursor", "rules")
		entries, err := os.ReadDir(rulesDir)
		if err != nil {
			return errors.New("cursor target artifact must include .cursor/rules/*.mdc")
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".mdc") {
				return nil
			}
		}
		return errors.New("cursor target artifact must include .cursor/rules/*.mdc")
	case "copilot":
		skillsDir := filepath.Join(targetDir, ".github", "skills")
		if exists(skillsDir) {
			return nil
		}
		if exists(filepath.Join(targetDir, "SKILL.md")) {
			return nil
		}
		return errors.New("copilot target artifact must include .github/skills/<skill-id>/ or SKILL.md")
	}
	return nil
}

func validateDestination(root, src, dest string) error {
	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return err
	}
	srcAbs, err := filepath.Abs(src)
	if err != nil {
		return err
	}
	destAbs, err := filepath.Abs(dest)
	if err != nil {
		return err
	}
	homeAbs, _ := filepath.Abs(userHome())
	agentxAbs := filepath.Join(rootAbs, ".agentx")

	if destAbs == string(filepath.Separator) || samePath(destAbs, rootAbs) || samePath(destAbs, homeAbs) {
		return errors.New("dangerous destination: refuses repository root, home directory, or filesystem root")
	}
	if withinPath(destAbs, rootAbs) {
		return errors.New("dangerous destination: refuses paths inside the AgentX repository")
	}
	if withinPath(destAbs, agentxAbs) {
		return errors.New("dangerous destination: refuses paths inside .agentx state")
	}
	if samePath(destAbs, srcAbs) || withinPath(destAbs, srcAbs) || withinPath(srcAbs, destAbs) {
		return errors.New("dangerous destination: refuses source path overlap")
	}
	if isSymlink(destAbs) {
		return errors.New("dangerous destination: destination must not be a symlink")
	}
	if info, err := os.Stat(destAbs); err == nil && !info.IsDir() {
		return errors.New("dangerous destination: destination exists and is not a directory")
	}
	return nil
}

func samePath(a, b string) bool {
	return filepath.Clean(a) == filepath.Clean(b)
}

func withinPath(path, parent string) bool {
	path = filepath.Clean(path)
	parent = filepath.Clean(parent)
	if samePath(path, parent) {
		return true
	}
	rel, err := filepath.Rel(parent, path)
	if err != nil {
		return false
	}
	return rel != "." && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

func copyTree(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Type()&os.ModeSymlink != 0 {
			rel, _ := filepath.Rel(src, path)
			return errors.New("copy refuses symlink: " + filepath.ToSlash(rel))
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		out := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(out, 0o755)
		}
		if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		file, err := os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, in)
		return err
	})
}

func hashTree(root string) (map[string]string, error) {
	hashes := map[string]string{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		sum := sha256.Sum256(data)
		rel, _ := filepath.Rel(root, path)
		hashes[filepath.ToSlash(rel)] = hex.EncodeToString(sum[:])
		return nil
	})
	return hashes, err
}

func flagValue(args []string, name string) string {
	for i, arg := range args {
		if arg == name && i+1 < len(args) {
			return args[i+1]
		}
	}
	return ""
}

func hasFlag(args []string, name string) bool {
	for _, arg := range args {
		if arg == name {
			return true
		}
	}
	return false
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func readText(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}

func printJSON(v any) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}

func userHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "~"
	}
	return home
}
