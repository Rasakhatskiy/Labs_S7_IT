<script>
	export let data;
</script>

<!-- header -->
<div class="sticky top-0 flex flex-col justify-center overflow-hidden bg-gray-50">
	<div class="flex items-center justify-between bg-blue-500 p-3">
		<a
			href="/"
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
			Database {data.dbname}
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
									<tr class="relative hover:bg-gray-100 dark:hover:bg-gray-700">
										<td
											class="py-7 px-6 text-sm  font-medium text-center text-gray-900 whitespace-nowrap dark:text-white"
										>
											<a
												class="absolute left-0 top-0 flex items-center justify-center w-full h-full"
												href="/databases/{data.dbname}/{table}">{table}</a
											>
										</td>
									</tr>
								{/each}
								<!-- 404 test -->
								<!-- <tr class="relative hover:bg-gray-100 dark:hover:bg-gray-700">
									<td
										class="py-7 px-6 text-sm  font-medium text-center text-gray-900 whitespace-nowrap dark:text-white"
									>
										<a
											class="absolute left-0 top-0 flex items-center justify-center w-full h-full"
											href="/databases/{data.dbname}/doesntexist">😭</a
										>
									</td>
								</tr> -->
								<!-- 404 test -->
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
