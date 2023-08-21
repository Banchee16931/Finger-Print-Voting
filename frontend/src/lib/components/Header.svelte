<script lang="ts">
    import logo from "../images/crown.png"
    import {page} from '$app/stores'
    let pageName = $page.url.pathname.substr($page.url.pathname.lastIndexOf('/'))
    $: noLogin = pageName == "/login"
</script>


<div class="navbar no-format-link">
    <div class="navbar spaced-container">
        <div class="brand hoverable">
            <a href="/"><img class="logo" src={logo} alt="Crown Logo"/><span class="text">GOV.UK</span></a>
        </div>
        {#if noLogin == false}
            <a href="/login" class="item hoverable">
                Login
            </a>
        {/if}
    </div>
</div>