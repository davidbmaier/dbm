# Database Maier

Monorepo for both front- and backend for the **D**ata**b**ase **M**aier project.

The production site is private, but I thought I'd publish the code anyway.

## Frontend

- Svelte/SvelteKit
- component library: flowbite-svelte

## Backend

- Go web server
- Postgres for structured data storage
- image files are stored on disk

## TODO

- pagination improvements (changeable page size, first/last buttons, label for how many total there are and which ones are currently shown, double)
- use placeholders for WorkTiles to prevent scroll position getting lost
- change the page title to reflect where you are
- add backend sorting support for works and artists
- add more filters for works and artists (e.g. extend works search to filter by artist name)
- add artist/artwork of the day for the home page
- clean up database entries (see script TODO comments)
- make header sticky
- add health check on login page to immediately redirect to /works if already logged in
- add more types of data (image descriptions, exhibitions, articles, etc.)
- timeline on artist page (birth year - death year with works in the middle)
