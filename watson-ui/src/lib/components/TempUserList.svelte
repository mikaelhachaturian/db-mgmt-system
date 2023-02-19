<script lang="ts">
  import TempUserDetails from "./TempUserDetails.svelte";
  import {activeTempUsers} from "../../stores"

  export let loginUser: User;
  export let token: string;

  let showActiveUsers = $activeTempUsers.length > 0 ? true : false
  // let showActiveUsers = true
  // const test_list = [{"ID":1,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true}, {"ID":2,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":3,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":12,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":11,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":10,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":9,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":8,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":7,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":6,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":5,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true},{"ID":4,"CreatedAt":"2023-01-22T11:05:23.234606919Z","UpdatedAt":"2023-01-22T17:11:01.116395097Z","firstname":"Gal","lastname":"Regev","email":"g.regev@easysend.io","picture":"https://lh3.googleusercontent.com/a/AEdFTp5tbSKRVvlDNkO1idzEkZ691Em5SeSwIYRcFaoT=s96-c","duration":"35m","IsTempAccessActive":true}]

</script>

<div class="user-list">
  {#if showActiveUsers }
    {#each $activeTempUsers as user (user?.ID)}
        <TempUserDetails {token} {user} {loginUser}/>
    {/each}
  {:else}
    <p>No users currently...</p>
  {/if}
</div>

<style>
  .user-list {
    width: 100%;
    display: flex;
    flex-flow: row wrap;
    justify-content: flex-start;
    align-items: center;
    margin: 0 auto;
  }

</style>
