// import { error } from '@sveltejs/kit';
// import { PUBLIC_API_BASE_URL } from '$env/static/public'

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
  const databaseName = params.dbname

  return {
    databaseName
  }
  
  // throw error(404, 'Not found');
}
