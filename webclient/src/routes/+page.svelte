<!-- npm run dev -->
<script>
	export let data;

	import { PUBLIC_API_BASE_URL } from '$env/static/public';
	import { invalidateAll } from '$app/navigation';
	import axios from 'axios';
	import { each } from 'svelte/internal';

	let newName = '';
	function onDelete(i) {
		let url = `${PUBLIC_API_BASE_URL}/databases/${i}`;

		axios
			.delete(url)
			.then(function (response) {
				console.log(response);
			})
			.catch(function (error) {
				console.log(error);
			});

		invalidateAll();
	}

	function onCreate() {
		if (!newName) {
			return;
		}

		let url = `${PUBLIC_API_BASE_URL}/databases`;

		console.log(url);
		console.log(newName);

		axios
			.post(url, newName, {
				headers: {
					'Content-Type': 'application/json'
				}
			})
			.then(function (response) {
				console.log(response);
			})
			.catch(function (error) {
				console.log(error);
			});

		addData = [];

		invalidateAll();
	}
</script>

<div class="sticky top-0 flex flex-col justify-center overflow-hidden bg-gray-50">
	<div class="flex items-center justify-between bg-blue-500 py-7">
		<div class="absolute left-1/2 -translate-x-1/2 text-lg font-bold text-gray-100">
			😎 Database Management System 😎
		</div>
	</div>
</div>

<div class="container mx-auto flex items-center justify-center flex-col mt-10">
	<div class="min-w-[50%] w-full sm:w-auto">
		<div class="flex flex-col justify-center">
			<div class="overflow-x-auto shadow-md sm:rounded-lg">
				<div class="inline-block min-w-full align-middle">
					<div class="overflow-hidden">
						<table class="min-w-full divide-y divide-gray-200 table-fixed dark:divide-gray-700">
							<tbody
								class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700"
							>
								{#each data.databases as database}
									<tr class="relative hover:bg-gray-100 dark:hover:bg-gray-700">
										<td
											class="py-2  text-sm  font-medium text-center text-gray-900 whitespace-nowrap dark:text-white"
										>
											<a
												class="flex items-center justify-center w-full h-full"
												href="/databases/{database.name}">{database.name}</a
											>
										</td>
										<td class="py-2 px-2  text-gray-900 whitespace-nowrap dark:text-white">
											<button
												on:click|preventDefault={() => onDelete(database.name)}
												class="w-full inline-block bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
											>
												Delete
											</button>
										</td>
									</tr>
								{/each}
								<tr>
									<td class="py-2">
										<input
											type="text"
											id="db_name"
											bind:value={newName}
											class="text-xs rounded-lg   block w-full p-2.5 bg-gray-800 border-gray-600 placeholder-gray-500 text-white outline-none"
											placeholder="new_db_name"
											required
										/>
									</td>
									<td class="py-2 px-2">
										<button
											on:click={onCreate}
											class="px-2 w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 rounded"
										>
											Create new database
										</button>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<style>
</style>
