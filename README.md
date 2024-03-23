# Database Maier

Monorepo for both front- and backend for the **D**ata**b**ase **M**aier project.

The production site is private, but I thought I'd publish the code anyway.

This project is based on a private digital art collection, the website is effectively just a way to look through the data.

## Frontend

- Svelte/SvelteKit
- component library: flowbite-svelte

## Backend

- Go web server
- Postgres for structured data storage
- image files are stored on disk

## Screenshots

![image](https://github.com/davidbmaier/dbm/assets/17618532/275463fa-0cb7-4c3c-b2b9-341439ffe9ac)
![image](https://github.com/davidbmaier/dbm/assets/17618532/50ba9177-9ac7-4ea2-b14a-230826f65445)

## TODO

- add backend sorting support for works and artists
- add more filters for works and artists (e.g. extend works search to filter by artist name)
- add artist/artwork of the day for the home page
- add discovery mode (possibly as a replacement for the random buttons) which adds random buttons to work and artist pages
- identify duplicates and show them as one artwork with different images
- figure out a proper token strategy for renewing (maybe just write new JWT on request when expiry is soon?)
- clean up database entries (see script TODO comments)
- use placeholders for WorkTiles to prevent scroll position getting lost
- global search bar in the header (with different result types)?
- make header sticky
- add check on login page to immediately redirect to /works if already logged in
- add more types of data (image descriptions, exhibitions, articles, etc.)
- timeline on artist page (birth year - death year with works in the middle)
