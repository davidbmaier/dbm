# Database Maier

Monorepo for both front- and backend for the **D**ata**b**ase **M**aier project.

## Frontend

- written in JS
- framework not decided yet
  - React + Remix
  - Vue + ?
- component library?

## Backend

- language not decided yet
  - Go
  - JS
  - something else?
- Postgres for structured data storage
- binary data storage not decided yet
  - S3?
  - something else?
  - CDN could be cool but would be overkill

## General design decisions

- auth
  - via a third-party OAuth provider
  - no registration, closed system - account setup is done manually by me
    - if there's a way to also serve something publicly, it would be even nicer of course
    - maybe just don't serve the actual works, but make all the other content available?
      - might be problematic too, since most of the content was originally published in books and magazines
- structure not decided yet
  - what would be the most natural way to present all the data? works, artists, literature?
  - a lot of crosslinking, maybe also to Wikipedia or other sources
    - this might require a bunch of automated scraping of the sources
