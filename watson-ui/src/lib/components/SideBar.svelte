<script>
    export let currentSelect;
</script>

<section id="sideNav">
    <ul>
        <li class:active="{currentSelect === '/db-access/user-list'}" on:click={()=> {currentSelect = "/db-access/user-list"}} on:keypress><a href="/db-access/user-list"><i class="fas fa-user"></i></a></li>
        <li class:active="{currentSelect === '/db-access'}" on:click={()=> {currentSelect = "/db-access"}} on:keypress><a href="/db-access"><i class="fas fa-database"></i></a></li>
    </ul>
</section>

<style>

    #sideNav {
        width: 7em;
        height: 100vh;
        background-color: rgb(255, 255, 255);
        border-right: 1px solid rgb(243, 243, 243);
        box-shadow: 6px 0px 3px -5px rgba(0,0,0,0.1);
    }

    ul {
        list-style-type: none;
        margin: 1em auto;
        padding: 0;
        width: 100%;
        background-color: #ffffff;
        text-align: center;
    }

    li a {
        display: block;
        color: #000;
        padding: 2em 0;
        text-decoration: none;
        transition: all .3s;
    }

    li a:hover, .active {
        background-color: rgb(243, 243, 243);
        border-left: 5px solid rgb(176, 36, 194);
    }

    li a:hover i {
        color: rgb(163, 163, 163);
    }

    i {
        color: rgb(121, 121, 121);
    }

</style>