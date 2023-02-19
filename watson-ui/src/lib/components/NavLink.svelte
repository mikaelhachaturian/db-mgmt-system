<script>
  import { page } from '$app/stores';
  export let title = "title";
  export let href = "/";
</script>

<a {href} class:active={$page.url.pathname === href}>{title}</a>

<style>
  a.active {
    background-color: rgba(0,0,0,0.24);
  }
</style>