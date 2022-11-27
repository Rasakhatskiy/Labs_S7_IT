import { error } from '@sveltejs/kit';
import { PUBLIC_API_BASE_URL } from '$env/static/public';
import axios from 'axios';

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	const dbname = params.dbname;
	console.log(dbname);

	try {
		const info = (await axios.get(`${PUBLIC_API_BASE_URL}/databases/${dbname}/join_tables`)).data;
		console.log(info);
		return {
			dbname,
			info
		};
	} catch {
		throw error(404, 'Not found');
	}

	// throw error(404, 'Not found');
}
