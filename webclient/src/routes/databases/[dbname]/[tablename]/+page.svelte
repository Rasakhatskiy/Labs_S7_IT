<script>
	/** @type {import('./$types').PageData} */
	export let data;
</script>

<!-- header -->
<div class="sticky top-0 flex flex-col justify-center overflow-hidden bg-gray-50">
	<div class="flex items-center justify-between bg-blue-500 p-3">
		<a
			href="/databases/{data.dbname}"
			class="flex items-center space-x-2 rounded bg-gray-100 py-1 px-2 text-slate-500 shadow-md hover:bg-white"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				stroke="currentColor"
				class="h-4 w-4"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M19.5 12h-15m0 0l6.75 6.75M4.5 12l6.75-6.75"
				/>
			</svg>
			<span>Back</span>
		</a>
		<div class="absolute left-1/2 -translate-x-1/2 text-lg font-bold text-gray-100">
			{data.dbname}/{data.tablename}
		</div>
	</div>
</div>

<!-- component -->
<!-- This is an example component -->
<div class="max-w-2xl mx-auto mt-10">
	<div class="relative shadow-md sm:rounded-lg">
		<table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
			<thead
				class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 "
			>
				<tr>
					{#each data.table.headers as header}
						<th scope="col" class="px-6 py-3">
							<div>{header.name}</div>
							<br />
							<div class="text-smol">
								{header.type}
							</div>
						</th>
					{/each}
					<th scope="col" class="px-6 py-3">
						<span class="sr-only">Edit</span>
					</th>
				</tr>

				<tr class="bg-gray-700 border-t-2 border-gray-600">
					{#each data.table.headers as header}
						<th scope="col" class="px-6 py-3">
							<input
								type="text"
								id="last_name"
								class="text-xs rounded-lg   block w-full p-2.5 bg-gray-800 border-gray-600 placeholder-gray-500 text-white outline-none"
								placeholder={header.name}
								required
							/>
						</th>
					{/each}
					<td class="px-6 py-4 text-right">
						<a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Add</a>
					</td>
				</tr>
			</thead>
			<tbody>
				{#each data.table.values as row}
					<tr
						class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
					>
						{#each row as value}
							<td class="px-6 py-4"> {value} </td>
						{/each}
						<td class="px-6 py-4 text-right">
							<a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
								>Edit</a
							>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>

	<script src="https://unpkg.com/flowbite@1.3.4/dist/flowbite.js"></script>
</div>

<style>
	.text-smol {
		font-size: 0.6rem;
		line-height: 1rem;
	}
</style>
