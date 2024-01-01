import { PUBLIC_API_URL } from '$env/static/public';
import { goto } from '$app/navigation';
import type { ArtistsResponse, ErrorResponse, ParsedResponse, WorksResponse } from '../types';

const sendRequest = async (
	url: string,
	method = `GET`,
	body: object | null = null
): Promise<ParsedResponse> => {
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
};

export const isErrorResponse = (input: object): input is ErrorResponse => {
	return (input as ErrorResponse).error !== undefined;
};

export const login = async (user: string, password: string) => {
	const response = await sendRequest(`${PUBLIC_API_URL}/api/login`, 'POST', { user, password });
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	} else {
		goto('/works');
	}
};

export const getArtists = async (
	search = '',
	page = 1
): Promise<ArtistsResponse | ErrorResponse> => {
	const response = await sendRequest(`${PUBLIC_API_URL}/api/artists?page=${page}&search=${search}`);
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	} else {
		return response.parsedBody as ArtistsResponse;
	}
};

export const getWorks = async (search = '', page = 1): Promise<WorksResponse | ErrorResponse> => {
	const response = await sendRequest(`${PUBLIC_API_URL}/api/works?page=${page}&search=${search}`);
	if (isErrorResponse(response.parsedBody)) {
		return response.parsedBody as ErrorResponse;
	} else {
		return response.parsedBody as WorksResponse;
	}
};