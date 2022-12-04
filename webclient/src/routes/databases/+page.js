import { error } from '@sveltejs/kit';
import { PUBLIC_API_BASE_URL } from '$env/static/public';
import axios from 'axios';

/** @type {import('./$types').PageLoad} */
export async function load() {
	try {
		const databases = (await axios.get(PUBLIC_API_BASE_URL + '/databases/')).data;
		console.log(databases);
		return {
			databases
		};
	} catch {
		throw error(404, 'Not found');
	}
}
