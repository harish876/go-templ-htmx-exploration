# TL;DR

1. Building out UI components using Go and HTMX with a touch of Client-side javascript.
2. Using Alpine JS, Sweet Alerts, and Vanilla JS for client-side interactivity.

# Goals

1. Make a Frontend App whose state is driven by the server and uses HTMX for data fetching.
2. Building out cool type-safe components using gol goodness.
3. Presenting an Application Scaffolding that is scalable and can be used for larger projects.

# Todos

1. Replace in-memory data store with turso.
2. Add loading states to data grids and iron out form components.
3. Build out other cool components, infinite scroll feature next.

# Instructions

1. npm run start: I know using npm sucks for a Go project, but I sometimes work out of a windows laptop (sigh) so it's a pain using make. Plus I don't know how to run concurrent processes using make. This command builds your tailwind, generates the template files using Go Templ, and runs air, so great HMR setup.
