<script lang="ts">
	import { getWorks, isErrorResponse } from "$lib/requests";
	import { error } from "@sveltejs/kit";
    import type {WorksResponse} from "../../types"

    export let page = 1;
    export let search = "";

    let worksData: WorksResponse;

    const fetchWorksData = async () => {
        const fetchedData = await getWorks(search, page);
        if (isErrorResponse(fetchedData)) {
            error(fetchedData.status, { message: fetchedData.error})
        } else {
            worksData = fetchedData as WorksResponse
        }
    }

    fetchWorksData();
</script>

<div>
    {JSON.stringify(worksData)}
</div>
