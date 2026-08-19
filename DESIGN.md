---
name: SABIFY
description: AI-powered two-sided education platform for teachers and students
colors:
  primary: "#4f46e5"
  primary-deep: "#3730a3"
  primary-wash: "#eef2ff"
  surface: "#ffffff"
  canvas: "#f8fafc"
  text: "#0f172a"
  text-secondary: "#475569"
  text-muted: "#64748b"
  text-light: "#94a3b8"
  border: "#e2e8f0"
  border-subtle: "#f1f5f9"
  success: "#16a34a"
  success-wash: "#dcfce7"
  warning: "#f59e0b"
  warning-wash: "#fef3c7"
  danger: "#dc2626"
  danger-wash: "#fee2e2"
  ai: "#7c3aed"
  ai-deep: "#4f46e5"
  ai-wash: "#f5f3ff"
typography:
  display:
    fontFamily: "Inter, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "clamp(3rem, 7vw, 5.5rem)"
    fontWeight: 800
    lineHeight: 0.98
    letterSpacing: "-0.055em"
  headline:
    fontFamily: "Inter, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "3rem"
    fontWeight: 700
    lineHeight: 1.1
  title:
    fontFamily: "Inter, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "1.5rem"
    fontWeight: 600
    lineHeight: 1.25
  body:
    fontFamily: "Inter, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "1rem"
    fontWeight: 400
    lineHeight: 1.5
  label:
    fontFamily: "Inter, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "0.75rem"
    fontWeight: 600
    lineHeight: 1.5
    letterSpacing: "0.04em"
rounded:
  sm: "8px"
  md: "12px"
  lg: "16px"
  xl: "20px"
  pill: "9999px"
spacing:
  xs: "4px"
  sm: "8px"
  md: "12px"
  base: "16px"
  lg: "24px"
  xl: "32px"
  2xl: "48px"
  3xl: "64px"
  4xl: "96px"
  section: "96px"
components:
  button-primary:
    backgroundColor: "{colors.primary}"
    textColor: "#ffffff"
    rounded: "{rounded.md}"
    padding: "0 22px"
    height: "50px"
  button-primary-hover:
    backgroundColor: "{colors.primary-deep}"
    textColor: "#ffffff"
    rounded: "{rounded.md}"
    padding: "0 22px"
    height: "50px"
  button-secondary:
    backgroundColor: "rgba(255, 255, 255, 0.75)"
    textColor: "{colors.text}"
    rounded: "{rounded.md}"
    padding: "0 20px"
    height: "50px"
  nav-link:
    backgroundColor: "transparent"
    textColor: "{colors.text-secondary}"
    rounded: "{rounded.sm}"
    padding: "8px 12px"
  card:
    backgroundColor: "{colors.surface}"
    textColor: "{colors.text}"
    rounded: "{rounded.md}"
    padding: "16px"
  badge:
    backgroundColor: "{colors.primary-wash}"
    textColor: "{colors.primary}"
    rounded: "{rounded.pill}"
    padding: "7px 12px"
---

# SABIFY Design System

## Overview

**Creative North Star: "Brutalist clarity"**

SABIFY is a two-sided education platform where teachers create and manage courses, quizzes, and assignments, while students engage with AI-personalized learning paths. The design system is built on a single conviction: structure is visible. Grid lines aren't hidden, they're part of the aesthetic. Every element exists because it does something. No decorative frills.

The palette is electric and vibrant: an indigo-to-violet gradient (`#4f46e5` to `#7c3aed`) is the signature accent that appears on every CTA, the brand mark, AI badges, and progress bars. This high-energy gradient sits against a calm, neutral slate base (`#f8fafc` background, `#ffffff` surfaces), creating the core tension: professional foundation, electric accents. The brand voice is a friendly mentor -- supportive without being patronizing, clear without being cold.

Typography uses Inter at a clean 400/500/600/700/800 weight scale. The hero headline uses a dramatic `clamp(3rem, 7vw, 5.5rem)` display size with tight 0.98 line-height and -0.055em tracking for impact. Body text stays at a comfortable 1.5 line-height with max-width 70ch for readability. The system uses a consistent 4px base spacing unit with scale steps at 4, 8, 12, 16, 24, 32, 48, 64, 80, 96.

Motion is purposeful: `fade-up` (600ms ease, translateY 24px) for scroll reveals, `float` (5s infinite) for the hero badge, and button hover transforms (`translateY(-2px)` with shadow increase at 200ms). All animations respect `prefers-reduced-motion: reduce`.

**Key Characteristics:**
- Visible grid structure with generous whitespace (96px section padding)
- Indigo-violet gradient as the single accent signature
- Glassmorphic sticky navbar (backdrop-filter blur 16px, semi-transparent white)
- Dashboard preview hero proving the product is real
- Clean Inter typography with dramatic display scaling
- Consistent border-radius vocabulary: 8px small, 12px cards, 16px panels, pill badges

## Colors

The palette is built on an indigo-violet accent against a neutral slate base. The contrast creates the feeling: professional foundation, electric accents.

### Primary
- **Indigo 600** (#4f46e5): The primary accent. CTAs, active nav states, links, brand mark gradient start, progress bar fills, AI badge backgrounds. Used sparingly but consistently.
- **Indigo 800** (#3730a3): Primary hover state. Darkens the accent on interaction.
- **Indigo 50** (#eef2ff): Primary wash. Hover backgrounds for nav links, subtle highlight panels, hero gradient start.

### Secondary
- **Violet 600** (#7c3aed): AI accent. Gradient partner to indigo (appears in `--gradient-ai`), AI insight card backgrounds, brand mark gradient end. Signals "AI is happening here."
- **Violet 50** (#f5f3ff): AI wash. AI feature section backgrounds, hero gradient midpoint.

### Neutral
- **Slate 900** (#0f172a): Text primary. Headings, body text, strong emphasis.
- **Slate 600** (#475569): Text secondary. Paragraphs, descriptions, supporting copy.
- **Slate 500** (#64748b): Text muted. Labels, metadata, timestamps, section labels.
- **Slate 400** (#94a3b8): Text light. Placeholders, disabled states.
- **Slate 200** (#e2e8f0): Border. Card borders, dividers, input borders.
- **Slate 100** (#f1f5f9): Border subtle. Lighter separators, topbar backgrounds.
- **Slate 50** (#f8fafc): Canvas. Page background.
- **White** (#ffffff): Surface. Cards, modals, panels, navbar background.

### Named Rules

**The Gradient Rule.** The indigo-to-violet gradient is the only gradient in the system. It appears on primary CTAs, the brand mark, AI badges, and progress bars. No other gradient exists. Its rarity is the point.

**The 10% Accent Rule.** The primary accent covers no more than 10% of any given screen. It guides the eye; it doesn't shout.

## Typography

**Display Font:** Inter (with system fallback stack: -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif)
**Body Font:** Inter (same stack)
**Label Font:** Inter (same stack, at smaller sizes with uppercase transforms)

**Character:** Clean, geometric, professional. The single-family approach creates consistency; weight variation (400-800) handles hierarchy. No decorative fonts. No serif. The type system is invisible when it works and precise when you look closely.

### Hierarchy
- **Display** (800, `clamp(3rem, 7vw, 5.5rem)`, 0.98 line-height): Hero headlines only. Tight line-height creates dramatic impact.
- **Headline** (700, 3rem/2.25rem, 1.1 line-height): Section headings (h1/h2). Bold, tight tracking.
- **Title** (600, 1.5rem, 1.25 line-height): Card titles, feature headings (h3). Semibold for emphasis without weight.
- **Body** (400, 1rem, 1.5 line-height): Paragraphs, descriptions. Max-width 70ch for readability.
- **Label** (600, 0.75rem, 0.04em letter-spacing, uppercase): Section labels, badges, metadata. Small but authoritative.

### Named Rules
**The Display Scale Rule.** Hero text uses `clamp(3rem, 7vw, 5.5rem)` -- it scales fluidly between mobile and desktop. Never set a fixed hero size below 3rem.

## Layout

Single-column centered layout with max-width 1280px. Container uses `min(100% - 2*page-padding, 1280px)` with auto margins for centering.

**Section rhythm:** 96px vertical padding between major sections. This generous spacing is deliberate -- it gives each section room to breathe and creates visual hierarchy through whitespace.

**Container padding:** 32px desktop, 24px tablet, 16px mobile. Responsive breakpoints handled in `responsive.css`.

**Dashboard preview:** 2-column grid (180px sidebar + flex-1 main) inside the hero section. Sidebar collapses to hamburger on mobile.

**Section headers:** max-width 700px, margin-bottom 48px. Section labels are uppercase, 14px semibold, indigo text, 0.08em letter-spacing.

## Elevation & Depth

The system uses a hybrid approach: shadows for structural depth, glassmorphism for atmospheric depth.

### Shadow Vocabulary
- **Subtle lift** (`0 2px 8px rgba(15,23,42,0.04)`): Dropdowns, small floating elements.
- **Card** (`0 8px 24px rgba(15,23,42,0.08)`): Cards, floating badges, default elevation.
- **Hero preview** (`0 30px 80px rgba(15,23,42,0.12), 0 8px 30px rgba(15,23,42,0.05)`): The dashboard preview in the hero. Dual shadow for depth + glow.
- **Modal** (`0 24px 60px rgba(15,23,42,0.12)`): Modal overlays, highest structural elevation.

### Atmospheric Depth
The navbar uses `backdrop-filter: blur(16px)` with `rgba(255, 255, 255, 0.88)` background for a frosted-glass effect. The hero section uses radial gradient glows (`rgba(99, 102, 241, 0.10)` centered top) and a 70%-width blur glow beneath the dashboard preview.

### Named Rules
**The Flat-By-Default Rule.** Surfaces are flat at rest. Shadows appear only as a response to state (hover, elevation, focus) or structural hierarchy (hero preview, modals).

## Shapes

The form language is soft but structured. Border-radius follows a consistent scale:

- **8px (sm):** Small interactive elements -- nav links, sidebar items, input fields.
- **12px (md):** Standard containers -- cards, buttons, stat panels, AI insight cards.
- **16px (lg):** Larger containers -- feature panels, dashboard analysis sections.
- **20px (xl):** Hero dashboard preview card.
- **Pill (9999px):** Badges, labels, hero label, performance bar tracks, progress indicators.

**Borders:** 1px solid throughout. `#e2e8f0` for standard borders (cards, inputs), `#f1f5f9` for subtle separators (topbar, sidebar dividers). No 2px or 3px borders anywhere.

**Clipping:** None. No overflow hidden on content containers. The hero dashboard preview uses `overflow: hidden` only to clip its internal content to the 20px radius.

## Components

### Buttons

Buttons are the workhorses of the system. Three variants, each with a clear job.

- **Shape:** 12px radius (md), 50px height (hero) or 40px (navbar). Transition: transform 200ms ease, box-shadow 200ms ease.
- **Primary:** Indigo-to-violet gradient background, white text, 14px semibold. Shadow: `0 8px 25px rgba(79, 70, 229, 0.22)`. Hover: `translateY(-2px)`, shadow increases to `0 12px 30px rgba(79, 70, 229, 0.30)`. Arrow icon slides right 4px on hover.
- **Secondary:** Semi-transparent white background (`rgba(255, 255, 255, 0.75)`), slate border, dark text. Hover: border shifts to indigo (`rgba(99, 102, 241, 0.30)`), background gains indigo tint, `translateY(-2px)`.
- **Ghost (nav):** Transparent background, rounded 8px. Hover: indigo background wash, text darkens. No transform on hover.

### Navigation

Sticky glassmorphic navbar. 76px min-height. Brand mark (34x34 gradient square with 2x2 dot grid) on the left, centered nav links, actions on the right.

- **Default:** White text, 14px medium weight. Transparent background.
- **Hover:** Text darkens, gains `rgba(238, 242, 255, 1)` background wash.
- **Mobile:** Hamburger button (42x42) with CSS-animated X transform. Slide-down panel with max-height transition (350ms ease).

### Cards

White background, 1px `#f1f5f9` border, 12px radius. Padding varies: 16px (stat cards), 18px (panels), 25px (dashboard main).

- **Stat cards:** 3-column grid, 12px gap. Value: 24px bold, -0.04em tracking. Change indicator: 11px green.
- **AI insight card:** Gradient background (`rgba(99, 102, 241, 0.06)` to `rgba(139, 92, 246, 0.03)`), indigo border (`rgba(99, 102, 241, 0.16)`), 10px radius.

### Hero Label

Pill-shaped badge above the hero title. Indigo border (`rgba(99, 102, 241, 0.18)`), indigo background wash (`rgba(99, 102, 241, 0.06)`), indigo text, 12px semibold uppercase. Includes a 7px dot with a 4px ring glow.

### Performance Bars

6px height, pill-shaped track (`#eef2f7`), gradient fill (indigo-to-violet). Layout: 110px topic label | 1fr bar | 35px score. Used in the hero dashboard preview to show student mastery.

## Do's and Don'ts

### Do:
- **Do** use the indigo-to-violet gradient for all primary CTAs and AI-related badges.
- **Do** maintain 96px vertical padding between major sections.
- **Do** use `clamp()` for hero typography to ensure fluid scaling.
- **Do** keep the navbar glassmorphic (backdrop-filter blur + semi-transparent white).
- **Do** use the 8/12/16/20px radius scale consistently.
- **Do** respect `prefers-reduced-motion: reduce` -- all animations already handle this.

### Don't:
- **Don't** add gradients beyond the single indigo-to-violet signature.
- **Don't** use shadows on flat-at-rest surfaces -- shadows are for hover/elevation only.
- **Don't** set fixed hero text sizes below 3rem -- always use the `clamp()` scale.
- **Don't** use border widths greater than 1px.
- **Don't** introduce additional font families -- Inter handles all roles through weight variation.
- **Don't** put more than 10% accent color on any single screen.
