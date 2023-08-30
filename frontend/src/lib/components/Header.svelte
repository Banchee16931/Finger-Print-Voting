<script lang="ts">
    import logo from "../images/crown.png"
    import {page} from '$app/stores'
    import type { UserData } from "$lib/types"
	import { invalidateAll } from "$app/navigation";

    // Stores information that is shown on pop-up
    export let userDetails: UserData
    
    // Tracks what the last value of the path is
    $: pageName = $page.url.pathname.substr($page.url.pathname.lastIndexOf('/'))
    // Removes the login button when on the login page
    $: noLogin = pageName == "/login"

    // Makes it so the first letter and only the first letter is capatalised in a word
    function nameCapatilsation(name: string): string {
        if (name === undefined) {
            return ""
        }

        if (name.length == 0) {
            return name
        }
        return name.charAt(0).toUpperCase()+name.slice(1)
    }

    // Will remove all the user details from the store and reset the auth cookies
    async function  logout(e: Event) {
        await fetch("/api/delete-authorization", { method: "DELETE" })
        invalidateAll()
    }

    let timer: number

    let profileHover: boolean = false

    export function resetInactivityTimer() {
        if (timer) {
            clearTimeout(timer)
        }
        const timeout = 5 * 1000  /*secs*/ * 60 /*mins*/
        timer = setTimeout(logout, timeout)
    }
</script>


<div class="header no-format-link">
    <div class="header spaced-container">
        <div class="brand hoverable">
            <a href="/"><img class="logo" src={logo} alt="Crown Logo"/><span class="text">GOV.UK</span></a>
        </div>
        {#if userDetails.level !== null }
            <span class="hover-extend item" class:hoverable-force={profileHover}
            on:mouseenter={() => {profileHover = true}}
            on:mouseleave={() => {profileHover = false}}
            role="menu" tabindex={null}>
                <span class="hoverable name">
                    {nameCapatilsation(userDetails.first_name)} {nameCapatilsation(userDetails.last_name)}
                </span>
            </span>
            <div class="user-options" class:user-options-appear={profileHover}
                on:mouseenter={() => {profileHover = true}}
                on:mouseleave={() => {profileHover = false}}
                role="menu" tabindex={null}>
                <span class="content-left" style="grid-area:id-tag">Username: </span>
                <span class="content-right" style="grid-area:id">{userDetails.username}</span>

                <span class="content-left" style="grid-area:access-tag">Access Level: </span>
                <span class="content-right" style="grid-area:access">{nameCapatilsation(userDetails.level)}</span>

                <span class="divider"></span>
                <span class="content-left" style="grid-area:first-tag">First Name: </span>

                <span class="content-right" style="grid-area:first">{nameCapatilsation(userDetails.first_name)}</span>
                <span class="content-left" style="grid-area:last-tag">Last Name: </span>
                
                <span class="content-right" style="grid-area:last">{nameCapatilsation(userDetails.last_name)}</span>
                <button class="content-left hoverable logout" on:click={logout}>Logout</button>
            </div>
        {:else if noLogin === false}
            <a href="/login" class="item hoverable">
                Login
            </a>
        {/if}
    </div>
</div>