/** @type {import('./$types').PageLoad} */
import { PUBLIC_API_BASE_URL } from '$env/static/public';

export async function load({ params }) {
	const dbname = params.dbname;
	console.log(dbname);
	const tablename = params.tablename;
	console.log(tablename);

	const table = await (
		await fetch(`${PUBLIC_API_BASE_URL}/databases/${dbname}/${tablename}`)
	).json();
	console.log(table);

	return {
		dbname,
		table
	};
}
