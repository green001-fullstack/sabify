# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

Two co-primary audiences in an educational context:

- **Teachers** create courses, author or auto-generate quizzes, review submissions, and monitor student progress through analytics. Their job is designing learning experiences and measuring outcomes.
- **Students** enroll in courses, take quizzes, receive personalized coaching, join AI-matched study groups, and track their own learning progress. Their job is acquiring knowledge and skills efficiently.

Neither audience is secondary. The product is a two-sided platform where both roles drive adoption.

## Product Purpose

SABIFY is an AI-powered Learning Management System that closes the loop between teaching and learning. Where traditional LMS tools stop at content delivery and manual assessment, SABIFY automates the full cycle: teachers teach, AI assesses and analyzes, the system personalizes, students collaborate, and outcomes improve. The product exists to make the teach-assess-analyze-personalize-collaborate-improve loop automatic instead of manual.

Success means a teacher can upload material and get a complete adaptive learning pipeline running with minimal manual intervention, and students experience personalized paths that measurably improve their outcomes.

## Positioning

Most LMS products are content repositories with quiz tools. SABIFY's differentiator is the **closed AI loop**: every stage of the learning cycle feeds the next automatically. Course material generates assessments, assessments produce analytics, analytics drive personalization, personalization enables smart collaboration, and collaboration outcomes refine the next cycle. A competitor could copy any single feature, but not the full feedback loop as a unified system.

## Operating Context

- Teachers create courses, upload materials, and manage quizzes through a dashboard workflow.
- Students browse courses, take quizzes, receive AI-generated feedback, and join study groups.
- The system runs as a single Go web server with server-side rendered HTML (no SPA framework).
- PostgreSQL is the backing store for all user, course, quiz, submission, and group data.
- AI features (quiz generation, learning coach, study group matching) are planned but not yet implemented in the codebase.
- The application is currently in MVP / hackathon stage — auth, course creation, and the public landing page are functional. Most student and teacher dashboard handlers are stubs.

## Capabilities and Constraints

**Confirmed capabilities:**
- User registration and authentication with role-based access (teacher / student)
- Course creation and listing
- Full landing page with marketing content
- Session management with secure cookies
- Role-gated route groups (teacher-only and student-only endpoints)

**Planned but not yet implemented:**
- AI quiz generation from course material
- AI teacher assistant and analytics
- AI learning coach (personalized paths)
- AI study group matching
- Quiz submission and automatic grading
- Student dashboard with progress tracking
- Materials upload and management
- Study group creation and membership

**Technical constraints:**
- Server-side rendered HTML via Go templates (no frontend JavaScript framework)
- PostgreSQL 16 required (via Docker Compose for local development)
- Single binary deployment model
- No API layer — all communication is traditional form submissions and page navigations

## Brand Commitments

- Product name: **SABIFY**
- No logo, color palette, typography, or voice commitments established yet.

## Evidence on Hand

- Working Go codebase at `cmd/web/` with Chi router, pgx database layer, and Go template caching
- Public landing page at `ui/html/pages/home/index.html` (2404 lines, fully designed)
- Component library in `ui/html/components/` (navbar, footer, course-card, quiz-card, sidebar)
- Static CSS design system in `ui/static/css/` (15 files covering variables, layout, responsive, animations)
- PostgreSQL schema in `migrations/001_initial_schema.sql` with 8 tables (users, courses, materials, quizzes, questions, submissions, study_groups, study_group_members)
- No test files exist yet (`*_test.go` — zero)
- No DESIGN.md or PRODUCT.md existed before this file

## Product Principles

1. **The loop is the product.** Every feature should connect to and reinforce the teach-assess-analyze-personalize-collaborate-improve cycle. Standalone features that don't feed the loop are out of scope.
2. **Both sides must feel complete.** Teachers and students are co-primary. Neither role gets a degraded experience.
3. **AI augments, not replaces.** AI handles the mechanical heavy lifting (generating quizzes, matching groups, analyzing patterns), but teachers and students make the final decisions.
4. **Server-side simplicity.** Prefer server-rendered pages over client-side complexity. The architecture should stay deployable as a single binary.
5. **Ship the loop before polishing.** Getting the full cycle working end-to-end, even crudely, matters more than making any single stage perfect.

## Accessibility & Inclusion

No specific accessibility requirements established yet. The product serves an educational context, so standard web accessibility (WCAG 2.1 AA) should be treated as a baseline expectation for future work.
