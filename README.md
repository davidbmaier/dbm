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
