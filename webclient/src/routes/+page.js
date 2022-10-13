// import { error } from '@sveltejs/kit';
import { PUBLIC_API_BASE_URL } from '$env/static/public'

/** @type {import('./$types').PageLoad} */
export async function load() {
  const serverData = await (await fetch(PUBLIC_API_BASE_URL + '/databases')).json()

  return {
    databases: serverData.databases
  }
  
  // throw error(404, 'Not found');
}
