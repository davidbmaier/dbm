import type { NumericRange } from '@sveltejs/kit';

export type ParsedResponse = Response & { parsedBody: object };

export type ErrorResponse = {
	error: string;
	status: NumericRange<400, 599>;
};

export type Artist = {
	id: number;
	name: string;
	birthYear: number;
	deathYear: number;
	numberOfWorks: number;
};

export type ArtistsResponse = {
	artists: Artist[];
	total: number;
};

export type Work = {
	id: number;
	workID: string;
	artistID: number;
	artist: Artist;
	name: string;
	creationInfo: string;
	material: string;
	size: string;
	owner: string;
	source: string;
};

export type WorksResponse = {
	works: Work[];
	total: number;
};
