<script>
	/** @type {import('./$types').PageData} */
	export let data;

	import { PUBLIC_API_BASE_URL } from '$env/static/public';
	import { invalidateAll } from '$app/navigation';
	import axios from 'axios';
	import { each } from 'svelte/internal';

	let tableData = data.info

	let selected1 = 0,
		selected2 = 0;

		let c1=0, c2=0;

	$: columnSelect1 = tableData.tables[selected1].headers.map(
		(element) => `${element.name} :: ${element.type}`
	);
	$: columnSelect2 = tableData.tables[selected2].headers.map(
		(element) => `${element.name} :: ${element.type}`
	);

	const onclick = async () => {
		const url = new URL(`${window.location.origin}/databases/${data.dbname}/joined`)

		url.searchParams.append("t1", tableData.tables[selected1].name)
		url.searchParams.append("t2", tableData.tables[selected2].name)
		url.searchParams.append("c1", tableData.tables[selected1].headers[c1].name)
		url.searchParams.append("c2", tableData.tables[selected2].headers[c2].name)

		window.location.href = url
	}

</script>

<div class="container mx-auto flex items-center justify-center flex-col mt-10">
	<div class="min-w-[50%] w-full sm:w-auto">
		<div class="flex flex-col justify-center">
			<div class="overflow-x-auto shadow-md sm:rounded-lg">
				<div class="inline-block min-w-full align-middle">
					<label for="s1" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
						>Select first table</label
					>
					<select
						bind:value={selected1}
						id="s1"
						class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					>
						<!-- <option selected>Choose a table 1</option> -->
						{#each tableData.tables as table, i}
							<option value={i}>{table.name}</option>
						{/each}
					</select>

					<label for="s2" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
						>Select second table</label
					>
					<select
						bind:value={selected2}
						id="s2"
						class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					>
						<!-- <option selected>Choose a table 2</option> -->
						{#each tableData.tables as table, i}
							<option value={i}>{table.name}</option>
						{/each}
					</select>

					<label for="s21" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
						
						>Select first column</label
					>
					<select
						id="s21"
						bind:value={c1}
						class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					>
						{#each columnSelect1 as header, i}
							<option value={i}>{header}</option>
						{/each}
					</select>

					<label for="s22" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
						>Select second column</label
					>
					<select
						id="s22"
						bind:value={c2}
						class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					>
						{#each columnSelect2 as header, i}
							<option value={i}>{header}</option>
						{/each}
					</select>
					<br />
					<a class="w-full block" href="#">
						<button
							on:click={onclick}
							class="bg-purple-700 hover:bg-purple-900 text-white font-bold py-2 px-4 rounded w-full"
						>
							Join tables
						</button>
					</a>
				</div>
			</div>
		</div>
	</div>
</div>
