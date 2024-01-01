<script lang="ts">
	import { getArtists, isErrorResponse } from "$lib/requests";
    import type {ArtistsResponse, ErrorResponse} from "../../types"

    export let page = 1;
    export let search = "";

    let artistsData: ArtistsResponse;
    let error: ErrorResponse;

    const fetchArtistsData = async () => {
        const fetchedData = await getArtists(search, page);
        if (isErrorResponse(fetchedData)) {
            error = fetchedData;
        } else {
            artistsData = fetchedData as ArtistsResponse;
        }
    }

    fetchArtistsData();
</script>

<div>
    {#if error}
        <p>{error.error}</p>
    {:else}
        {JSON.stringify(artistsData)}
    {/if}
</div>
