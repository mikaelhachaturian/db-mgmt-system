<script lang="ts">
  export let itemLabel: string;
	export let highlighted: boolean;
</script>

<li class="autocomplete-items" class:autocomplete-active={highlighted} on:click on:keydown>{@html itemLabel}</li>

<style>
li.autocomplete-items {
	list-style: none;
  z-index: 99;
  /*position the autocomplete items to be the same width as the container:*/
  top: 100%;
  left: 0;
  right: 0;
	padding: 10px;
  cursor: pointer;
  background-color: #fff;
  border-radius: .5em;
  border-bottom: 1px solid #f2f2f2;
  transition: all .5s;
  color: #696969;
}

li.autocomplete-items:hover {
  /*when hovering an item:*/
  background-color: #f5f5f5;

}

li.autocomplete-items:active {
  /*when navigating through the items using the arrow keys:*/
  background-color: DodgerBlue !important;
  color: #ffffff;
}	
	
.autocomplete-active {
  /*when navigating through the items using the arrow keys:*/
  background-color: DodgerBlue !important;
  color: #ffffff;
}
</style>