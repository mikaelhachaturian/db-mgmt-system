<script lang="ts">
import { activeTempUsers } from "../../../stores";
import type { PageServerData } from './$types';
import TempUserList from "$lib/components/TempUserList.svelte";

export let data: PageServerData;
const { user } = data.session

activeTempUsers.set(data.users)

</script>

<section>
  <div class="page_body">
    <h3>Active Users in DEV DB</h3>
    <TempUserList token={data.token} loginUser={user}/>
  </div>
</section>

<style>

section {
    margin: 0 auto;
    width: 80%;
    display: block;
    overflow-y: scroll;
    height: 80vh;
  }

  ::-webkit-scrollbar {
    width: 10px;
  } 

::-webkit-scrollbar-track {
  border-radius: 1em;
  background: #ffffff;
  border: 1px solid #ededed;
}

::-webkit-scrollbar-thumb {
  background: rgb(103, 8, 120);
  border-radius: 1em;
}

::-webkit-scrollbar-thumb:hover {
  background: rgb(148, 10, 172);
}

  .page_body {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
  }
  h3 {
    color: rgb(96, 96, 96);
    padding: .5em 0;
    border-radius: .7em;
    margin: 2em 0;
  }

</style>