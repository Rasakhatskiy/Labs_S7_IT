// import { error } from '@sveltejs/kit';
import { PUBLIC_API_BASE_URL } from '$env/static/public';

/** @type {import('./$types').PageLoad} */
export async function load() {
	const databases = await (await fetch(PUBLIC_API_BASE_URL + '/databases')).json();
	console.log(databases);
	return {
		databases
	};

	// throw error(404, 'Not found');
}
