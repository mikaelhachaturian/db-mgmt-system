<script lang="ts">
	import CreateUserForm from "./CreateUserForm.svelte"
  import { activeTempUsers } from "../../stores";
  import type { PageServerData } from './$types';

  export let data: PageServerData;
  const { user } = data.session

  activeTempUsers.set(data.users)
</script>

<section>
  <div class="page_body">
    <h3>Create Temporary User for Dev Shared DB</h3>
    <CreateUserForm token={data.token} {user} dbs={data.dbs}/>
  </div>
</section>

<style>

  section {
    margin: 0 auto;
  }

  .page_body {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  h3 {
    color: rgb(46, 46, 46);
    padding: .7em 1em;
    border-radius: .7em;
    margin: 2em auto;
  }

</style>
