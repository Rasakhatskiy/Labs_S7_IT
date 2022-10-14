<script>
	export let data;
</script>

<!-- header -->
<div class="sticky top-0 flex flex-col justify-center overflow-hidden bg-gray-50">
	<div class="flex items-center justify-between bg-blue-500 p-3">
		<div class="flex items-center space-x-2 rounded bg-gray-100 py-1 px-2 text-slate-500 shadow-md">
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
			<span><a href="/">Back</a></span>
		</div>
		<div class="text-lg font-bold text-gray-100">Database {data.dbname}</div>
		<div class="flex items-center space-x-5 text-gray-100">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 24 24"
				fill="currentColor"
				class="h-6 w-6"
			>
				<path
					fill-rule="evenodd"
					d="M6.32 2.577a49.255 49.255 0 0111.36 0c1.497.174 2.57 1.46 2.57 2.93V21a.75.75 0 01-1.085.67L12 18.089l-7.165 3.583A.75.75 0 013.75 21V5.507c0-1.47 1.073-2.756 2.57-2.93z"
					clip-rule="evenodd"
				/>
			</svg>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 24 24"
				fill="currentColor"
				class="h-6 w-6"
			>
				<path
					fill-rule="evenodd"
					d="M5.625 1.5H9a3.75 3.75 0 013.75 3.75v1.875c0 1.036.84 1.875 1.875 1.875H16.5a3.75 3.75 0 013.75 3.75v7.875c0 1.035-.84 1.875-1.875 1.875H5.625a1.875 1.875 0 01-1.875-1.875V3.375c0-1.036.84-1.875 1.875-1.875zm6.905 9.97a.75.75 0 00-1.06 0l-3 3a.75.75 0 101.06 1.06l1.72-1.72V18a.75.75 0 001.5 0v-4.19l1.72 1.72a.75.75 0 101.06-1.06l-3-3z"
					clip-rule="evenodd"
				/>
				<path
					d="M14.25 5.25a5.23 5.23 0 00-1.279-3.434 9.768 9.768 0 016.963 6.963A5.23 5.23 0 0016.5 7.5h-1.875a.375.375 0 01-.375-.375V5.25z"
				/>
			</svg>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 24 24"
				fill="currentColor"
				class="h-6 w-6"
			>
				<path
					fill-rule="evenodd"
					d="M5.625 1.5H9a3.75 3.75 0 013.75 3.75v1.875c0 1.036.84 1.875 1.875 1.875H16.5a3.75 3.75 0 013.75 3.75v7.875c0 1.035-.84 1.875-1.875 1.875H5.625a1.875 1.875 0 01-1.875-1.875V3.375c0-1.036.84-1.875 1.875-1.875zm5.845 17.03a.75.75 0 001.06 0l3-3a.75.75 0 10-1.06-1.06l-1.72 1.72V12a.75.75 0 00-1.5 0v4.19l-1.72-1.72a.75.75 0 00-1.06 1.06l3 3z"
					clip-rule="evenodd"
				/>
				<path
					d="M14.25 5.25a5.23 5.23 0 00-1.279-3.434 9.768 9.768 0 016.963 6.963A5.23 5.23 0 0016.5 7.5h-1.875a.375.375 0 01-.375-.375V5.25z"
				/>
			</svg>
		</div>
	</div>
</div>

<!-- main body -->
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
								{#each data.tables as table}
									<tr class="hover:bg-gray-100 dark:hover:bg-gray-700">
										<td
											class="py-4 px-6 text-sm font-medium text-center text-gray-900 whitespace-nowrap dark:text-white"
										>
											{table}
										</td>
									</tr>
								{/each}
							</tbody>
						</table>
					</div>
				</div>
			</div>
			<button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
				Create new table
			</button>
		</div>
	</div>
</div>
