import { error } from '@sveltejs/kit';
import { PUBLIC_API_BASE_URL } from '$env/static/public';
import axios from 'axios';

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	const dbname = params.dbname;
	console.log(dbname);

	try {
		const tables = (await axios.get(PUBLIC_API_BASE_URL + '/databases/' + dbname)).data;
		console.log(tables);
		return {
			dbname,
			tables
		};
	} catch {
		throw error(404, 'Not found');
	}
}
