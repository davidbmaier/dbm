import { PUBLIC_API_URL } from '$env/static/public';
import { _ } from 'svelte-i18n';
import { get } from 'svelte/store';
import { goto } from '$app/navigation';
import type {
	Artist,
	ArtistsResponse,
	ErrorResponse,
	ParsedResponse,
	Work,
	WorksResponse
} from '../types';

const sendRequest = async (
	url: string,
	method = `GET`,
	body: object | null = null
): Promise<ParsedResponse | { parsedBody: { error: string } }> => {
	try {
		const options: RequestInit = {
			method: method,
			credentials: 'include'
		};
		if (body) {
			options.body = JSON.stringify(body);
			options.headers = {
				'Content-Type': 'application/json'
			};
		}
		const response = (await fetch(url, options)) as ParsedResponse;

		if (response.headers.get('Content-Type') === `application/json`) {
			response.parsedBody = { ...(await response.json()), status: response.status };
		}

		if (response.status === 401 || response.status === 403) {
			console.log('Authorization failed, redirecting to login');
			goto('/');
		}

		return response;
	} catch (error) {
		return {
			parsedBody: { error: get(_)('error.networkError') }
		};
	}
};

export const isErrorResponse = (input: object): input is ErrorResponse => {
	return (input as ErrorResponse)?.error !== undefined;
};

export const login = async (user: string, password: string) => {
	const response = await sendRequest(`${PUBLIC_API_URL}/api/login`, 'POST', {
		username: user,
		password
	});
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	} else {
		goto('/home');
	}
};

export const logout = async () => {
	await sendRequest(`${PUBLIC_API_URL}/api/logout`, 'POST');
	goto('/');
};

export const updatePassword = async (password: string) => {
	const response = await sendRequest(`${PUBLIC_API_URL}/api/users`, 'PATCH', { password });
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	}
};

export const getArtists = async (
	search = '',
	page = 1,
	pageSize = 20
): Promise<ArtistsResponse | ErrorResponse> => {
	const response = await sendRequest(
		`${PUBLIC_API_URL}/api/artists?page=${page}&pageSize=${pageSize}&search=${search}`
	);
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	} else {
		return response.parsedBody as ArtistsResponse;
	}
};

export const getArtist = async (id: number): Promise<Artist | ErrorResponse> => {
	const response = await sendRequest(`${PUBLIC_API_URL}/api/artists/${id}`);
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	} else {
		return response.parsedBody as Artist;
	}
};

export const getWorks = async (
	search = '',
	page = 1,
	pageSize = 20,
	artistID?: number
): Promise<WorksResponse | ErrorResponse> => {
	let url = `${PUBLIC_API_URL}/api/works?page=${page}&pageSize=${pageSize}&search=${search}`;
	if (artistID) {
		url += `&artistID=${artistID}`;
	}
	const response = await sendRequest(url);
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	} else {
		return response.parsedBody as WorksResponse;
	}
};

export const getWork = async (id: number): Promise<Work | ErrorResponse> => {
	const response = await sendRequest(`${PUBLIC_API_URL}/api/works/${id}`);
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	} else {
		return response.parsedBody as Work;
	}
};
