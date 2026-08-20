---
target: home page
total_score: 16
max_score: 32
na_heuristics: 7,10
p0_count: 2
p1_count: 2
timestamp: 2026-08-20T12-16-25Z
slug: ui-html-pages-home-index-html
---
# Design Critique: SABIFY Home Page

Method: dual-agent (A: design director review · B: detector + source scan). Browser visualization unavailable (no browser tool in this session); CLI detector ran on a reconstructed full-page replica for full rule coverage.

## Design Health Score

| # | Heuristic | Score | Key Issue |
|---|-----------|-------|-----------|
| 1 | Visibility of System Status | 1 | Only mock status exists; every real action dies silently or with a bare 500 |
| 2 | Match System / Real World | 4 | Genuinely excellent — concrete, correct physics/teaching language throughout |
| 3 | User Control and Freedom | 2 | Footer links 404 with no recovery; in-page nav is one-way only |
| 4 | Consistency and Standards | 2 | Self-consistent execution, but violates its own DESIGN.md spec + platform standards |
| 5 | Error Prevention | 1 | No validation UI; the one real form can't even load; broken anchors ship |
| 6 | Recognition Rather Than Recall | 3 | Consistent section-label pattern; ✦-marks-AI convention carried everywhere |
| 7 | Flexibility and Efficiency | n/a | Persuade surface; no expert/shortcut workflows to optimize |
| 8 | Aesthetic and Minimalist Design | 2 | Polished but not minimal: 5 full dashboard mocks + decorative orb/glows/pulse |
| 9 | Error Recovery | 1 | Broken links → bare 404; no styled 404, no guidance |
| 10 | Help and Documentation | n/a | No help needed on a static landing page |
| **Total** | | **16/32 (50%)** | **Acceptable/Poor boundary** |

Renormalized because heuristics 7 and 10 are n/a. The visual layer alone would score higher; every number is dragged down by a conversion funnel that 500s.

## Design Specificity Verdict

**Content is composed for this product; the visual skin is category-generic AI-SaaS.**

The proof lives in the words: the six-section loop (Teach/Assess/Analyze/Personalize/Collaborate/Improve) *is* the product loop from PRODUCT.md; the teacher/student tab split mirrors the two-co-primary positioning; and a single pedagogical thread — "force diagrams / Newton's Laws," the James/JD persona — recurs through the hero insight, teacher gap card, coach chat, roadmap, and study-group match, including the complementary pair (James weak in Momentum / Sarah strong in Momentum). That is genuine closed-loop thinking. But strip the copy and the surface is interchangeable: indigo-violet gradient chips, white cards, mock dashboards, ✦ icons, radial glows — the standard "AI for X" vocabulary. The differentiation is verbal, not visual.

**Deterministic scan:** 72 findings on the reconstructed full-page render (26 undersized-ui-text, 11 tiny-text, 7 kicker-above-heading, 7 nested-cards, 6 cramped-padding, 6 clipped-overflow-container, 2 low-contrast, plus single hits for gradient-text, dark-glow, skipped-heading, all-caps-body, overused-font, extreme-negative-tracking, gpt-thin-border-wide-shadow). False positives: all 5 `design-system-color` black findings (engine can't resolve `var(--color-text-primary)`), most nested-cards/cramped-padding (intentional mock previews). The detector caught what the human review would miss: the systematic <12px type floor breach, the 2.3:1 answer-key "C", the `overflow:hidden` clipping hazard on six sections, and the dead-link/500 funnel.

**Visual overlays:** none — no browser tool exposed this session; no user-visible overlay available. Fallback signal used: source + CLI evidence.

## Overall Impression

A genuinely well-argued landing page — the hero is the best moment, the force-diagrams through-line is real product thinking, and the AI-signaling grammar is legible — that is structurally betrayed by its own foundation. The navbar never renders, every register CTA 500s, the footer's trust links 404, and the design code contradicts its own committed spec on gradients, radii, borders, and display scale. The single biggest opportunity: fix the broken funnel first, then decide whether five full-width mock dashboards are proof or overload.

## What's Working

1. **The hero synthesis.** Copy, display type, dual CTA, and a mock dashboard whose data (Force & motion 86%, Force diagrams 42%) you can verify against the AI insight card in the same screenshot. The floating badge — "AI insight generated / Based on 124 student responses" — is specific enough to be believable.
2. **The narrative through-line.** "Force diagrams" is planted in the hero and resolved consistently across teacher panel, student coach, roadmap, and study-group match. The page demonstrates the closed loop it sells.
3. **The AI-signaling system.** ✦ mark, gradient chip, "AI recommendation/coach/match" labels — a coherent, repeated "AI is happening here" vocabulary. Plus `prefers-reduced-motion` is genuinely honored.

## Priority Issues

**P0 — Every conversion path 500s.**
What: The template cache globs only `./ui/html/pages/*/*.html`, yielding a single cached `index.html`. `register.html`/`login.html` live in `ui/html/auth/` (both 0 bytes) and are never parsed, so `render("register.html")` hits a missing cache key → HTTP 500. All six register CTAs plus navbar "Log in" die here.
Why: The persuasive argument is technically true; the conversion end of the funnel is broken. All landing-page effort ships into a void.
Fix: Parse `auth/*` and `layouts/*` into the cache; implement real register/login templates; add a test asserting GET /register returns 200. Command: `$impeccable harden`.

**P0 — Footer trust links are 404s.**
What: `/about`, `/contact`, `/privacy`, `/terms` have no routes — four guaranteed 404s on the one surface a skeptical educator checks before registering.
Why: Trust-destroying; a security-conscious visitor who clicks "Privacy" gets a bare 404 and leaves.
Fix: Add the four routes (even minimal placeholder pages) or remove the links until they exist. Command: `$impeccable harden`.

**P1 — The navbar never renders; no landmarks exist.**
What: `navbar.html`/`footer.html` are dead code — base.html renders only `{{ block "body" }}` and the page never invokes them. The rendered page has no `<header>`, no `<main>` wrapper, no nav landmark. The navbar links also point to `#for-teachers`/`#for-students` ids that don't exist.
Why: DESIGN.md's flagship glassmorphic component is absent; mobile users get zero top-level navigation on a 2404-line page; screen-reader landmark nav finds nothing.
Fix: Invoke the navbar in the layout, wrap content in `<main>`, add the missing anchor targets. Command: `$impeccable layout`.

**P1 — Role intent is dropped at the boundary.**
What: Six CTAs pass `?role=teacher|student`, but `showRegisterForm` never reads the query param — every role-specific promise resolves to an identical generic form.
Why: The page's entire two-sided architecture is sold on role; the first real screen discards it.
Fix: Read the param, persist it as a hidden field, prefill the role choice. Command: `$impeccable onboard`.

**P2 — Fake interactivity in a persuasive surface.**
What: "Generate quiz" and "Review question" are real `<button>`s; `.quiz-generate` runs a 1.4s fake spinner and returns to Ready with no output change. The dashboard mock's "Help" item looks clickable and isn't.
Why: A visitor who clicks a button and watches it lie has learned the product fakes it — poisoning every real claim on the page.
Fix: Make mock controls decorative (`div`/`aria-hidden`) or wire them to something real. Command: `$impeccable delight`.

**P2 — Design-spec violations compound.**
What: #2563eb (a blue not in the palette) inside the hero headline gradient; two variants of the signature gradient on avatar/chart-bar; a second dark gradient on the final CTA card; 14 distinct radius values vs. the committed 8/12/16/20; 2px/3px borders where the spec says 1px max; 800-weight titles vs committed 700; mobile hero at a fixed 2.35rem below the 3rem display floor; 8px footer note text.
Why: A design system only works if code honors it; "Brutalist clarity" reads as carelessness when the numbers disagree with the spec.
Fix: Normalize gradients to `--gradient-ai`, map radii onto the 8/12/16/20/pill scale, restore the display floor. Commands: `$impeccable colorize` + `$impeccable shape`.

**P3 — JS-dependent visibility.**
What: `.fade-up` (whole hero) and `.learning-step` default to `opacity: 0`, only becoming visible via IntersectionObserver. If JS fails or loads late, the hero and learning-loop are invisible.
Why: Catastrophic on flaky mobile and in security-hardened browsers — on a server-rendered page whose architecture promises no-JS resilience.
Fix: Add `<html class="no-js">` stripped by JS, and gate the opacity:0 states behind it. Command: `$impeccable harden`.

**P3 — Accessibility gaps in real controls.**
What: Quiz-builder `<label>`s have no `for`, inputs/selects no `id` (four unlabeled controls). The tablist implements click + aria-selected but no ArrowKey nav, no roving tabindex, no `aria-controls`; panels lack `role="tabpanel"`. No `:focus-visible` styles anywhere.
Why: WCAG 2.1 AA is the documented baseline; these are mechanical, standard failures.
Fix: Wire label/for pairs; implement the ARIA tabs pattern; add global focus-visible styling. Command: `$impeccable audit`.

**P3 — Mobile match-visual fragility.**
What: `.student-card` is fixed-width and absolutely positioned; at ≤767px the whole `.matching-area` is `scale(0.88)` — two absolute cards + center circle + rotated connection line + bottom chip on a 480px container is one line-wrap from overlapping.
Fix: Reflow to a stacked layout at ≤767px instead of scaling. Command: `$impeccable layout`.

## Persona Red Flags

**Jordan (confused first-timer):** No navbar — the trusted wayfinding layer of a first visit doesn't exist. "Get started" → HTTP 500; Jordan cannot register. The mock "Help" sidebar item is a fake affordance. "About"/"Contact" → 404. Two role-specific CTAs at the bottom both land on the same broken, role-blind form.

**Riley (stress tester):** Arrow keys do nothing in the intelligence switcher; focus lands on unlabeled selects; no visible focus on primary buttons. Disable JS → the hero and learning loop vanish (opacity:0). "Generate quiz" is a 1.4s lie — the card never changes. `?role=teacher` vs `?role=student` return identical responses. Labels at 10-11px in #94a3b8 sit at ≈2.9:1, failing WCAG AA. Every footer link dies (four 404s, one 500, two anchors to non-existent ids).

**Casey (distracted mobile user):** Mobile hero drops to a fixed 2.35rem — the headline loses its reason to exist on the primary mobile viewport. Five stacked full-width dashboard panels = endless thumb-scroll with no CTA until the final card. The sticky navbar Casey would rely on to jump back doesn't exist. Final CTA taps send Casey to a 500.

## Minor Observations

- `layouts/public.html` and `layouts/index.html` are dead files; base.html loads all four JS files on every page.
- Three different brand treatments in one page: "LI" dashboard mark, 2x2 dot grid navbar mark, ✦ footer mark.
- The same fake stat numbers (82/94/76) recur in hero and teacher panel — a sharp reader notices.
- Unicode glyph icons (◈ ▣ ✓ ◌ ✦ ◎) have no accessible names and render inconsistently across platforms.
- The quiz AI explanation (correct F=ma math) is a genuinely educational detail most AI-product pages get wrong.
- `html { font-size: 112.5% }` makes every px claim in DESIGN.md optically ~12.5% larger in the browser.

## Questions to Consider

1. If the only real action is register, and register is broken — what is the page actually converting? Before touching a gradient, is the move to cut this to a 2-screen page and pour the mocks into a real product tour after signup?
2. The emotional arc is entirely about *deficits* (gaps, struggles, needs practice). What if the story were a student who used to fail now leads the class — what would that do to the hero badge and mid-page mocks?
3. The spec says structure is visible and nothing is decorative — but the page ships a floating badge, pulsing connection, orbiting orb, four glow blobs, and six gradients the spec forbids. Is the design system wrong for a persuasion surface, or is the page wrong for the system?
