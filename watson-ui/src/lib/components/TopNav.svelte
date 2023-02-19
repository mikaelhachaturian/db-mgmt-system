<script lang="ts">
    export let currentUser: User;
    export let isLoggedIn: boolean;
    export let logoutAction: any;
</script>

<nav>
    <section class="lside">
        {#if isLoggedIn}
            <img id="avater" src={currentUser?.picture} alt="avater" referrerpolicy="no-referrer"/>
            <div class="dropdown">
                Welcome back! <span>{currentUser?.firstname} {currentUser?.lastname} <i class="fas fa-angle-down"></i></span>
                <div class="dropdown-content">
                    <ul>
                        <li><a on:click={logoutAction} href={'/'}> Logout</a></li>
                    </ul>
                </div>
            </div>
        {/if}
    </section>
    <a href="/db-access"><img id="logo" src="/watson-logo.png" alt="logo"/></a>
</nav>

<style>

    nav {
        width: 100vw;
        display: block;
        height: auto;
        background-color: #fff;
        padding: 1em 5em;
        box-shadow: 0px 2px 4px rgb(0 0 0 / 8%);
        display: flex;
        flex-direction: row-reverse;
        justify-content: space-between;
        align-items: center;
        z-index: 3;
    }

    .lside {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: max-content;
        height: 100%;
    }

    .dropdown > span {
        display: block;
        cursor: pointer;
        background-color: rgb(255, 255, 255);
        border-radius: 1em;
        padding: .3em 1em;
        margin-top: .3em;
        text-align: center;
        border-bottom: 2px solid purple;
    }

    .dropdown > span > i {
        padding-left: 5px;
    }

    #avater {
        width: 40px;
        height: auto;
        border-radius: 50%;
        overflow: hidden;
        margin-right: 2em;
    }

    #logo {
        width: 170px;
        height: auto;
        transition: all .3s;
    }

    #logo:hover {
        scale: 1.15;
    }

    .dropdown {
        position: relative;
        display: inline-block;
        height: 100%;
    }

    .dropdown-content {
        position: absolute;
        background-color: #ffffff;
        box-shadow: 5px 12px 3px -10px rgba(0,0,0,0.1);
        z-index: 1;
        right: 0px;
        top: 65px;
        border-radius: .3em;
        border-top: 3px solid #e5e5e5;
        transition: opacity .5s, visibility .5s;
        opacity: 0;
        visibility: hidden;
    }

    .dropdown-content ul {
        list-style: none;
    }

    .dropdown-content ul li {
        padding: 1em 1.8em;
    }

    .dropdown-content ul li:hover {
        text-decoration: underline;
        background-color: rgb(251, 251, 251);
    }

    .dropdown:hover .dropdown-content {
        visibility: visible;
        opacity: 1;
    }

    a {
        text-decoration: none;
    }

</style>