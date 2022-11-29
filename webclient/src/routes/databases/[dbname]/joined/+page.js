/** @type {import('./$types').PageLoad} */
import { PUBLIC_API_BASE_URL } from '$env/static/public';
import axios from 'axios';
import { error } from '@sveltejs/kit';

/** @type {import('./$types').PageLoad} */
export async function load({ params, url }) {
	const dbname = params.dbname;
	console.log(dbname);

	const fetchUrl = new URL(`${PUBLIC_API_BASE_URL}/databases/${dbname}/joined_tables`);

	fetchUrl.searchParams.append('t1', url.searchParams.get('t1'));
	fetchUrl.searchParams.append('t2', url.searchParams.get('t2'));
	fetchUrl.searchParams.append('c1', url.searchParams.get('c1'));
	fetchUrl.searchParams.append('c2', url.searchParams.get('c2'));

	try {
		const table = (await axios.get(fetchUrl)).data;
		return {
			dbname,
			table
		};
	} catch {
		throw error(404, 'Not found');
	}
}
