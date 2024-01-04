# Database Maier

Monorepo for both front- and backend for the **D**ata**b**ase **M**aier project.

## Frontend

- Svelte/SvelteKit
- component library: flowbite-svelte

## Backend

- Go web server
- Postgres for structured data storage
- images are stored on disk

## Data structure

- Overview
  - index.txt: contains all cataloged entries with references to works in Gallery
  - indexIncomplete.txt: contains uncataloged entries that should reference works in Gallery/Uncataloged (IDs may not be accurate, references might be missing, entries may not cover all uncataloged works)
- Gallery
  - A-Z: Works that have mostly been cataloged by the artist's last name
  - Uncataloged: Works referencing the artist's last name and that should largely be referenced in Overview/indexIncomplete.txt
- Texts about Works
  - Files: Documents referencing works in Gallery/A-Z
  - Uncataloged: Documents referencing works somewhere in Gallery (IDs may not be accurate)
- Critiques (of exhibitions, books and other meta-works)
  - index.doc: Overview of all critiques and their metadata
  - Files: Documents without further references
- Exhibitions
  - index.doc: Overview of all exhibitions and their metadata
  - Files: Documents without further references
- Biographies
  - Files: Documents without further references (but file names may correspond with that artist's works' IDs)
- Articles
  - index.doc: Overview of all articles and their metadata
  - Files: Documents without further references
- Bibliography
  - A-Z: Documents without further references (that have been used as sources for texts, critiques, exhibitions, biographies and articles)

## TODO

- extend works search to cover artist names
- identify works with missing files and fix them/remove them from the database
- improve styling for work pages
- add reset button to search input
- smooth loading of images
- add footer
- make header sticky
- logout button
- add backend sorting support for works and artists
- add more filters for works and artists
- add health check on login page to immediately redirect to /works if already logged in
