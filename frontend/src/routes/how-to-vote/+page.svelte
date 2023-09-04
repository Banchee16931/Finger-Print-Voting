<script>
	import Hero from "$lib/components/Hero.svelte";

   
    import register from "$lib/images/register.png"
    import registerForm from "$lib/images/register form.png"
    import login from "$lib/images/login.png"
    import loginForm from "$lib/images/login-form.png"


    import vote from "$lib/images/vote.png"
    import votePage from "$lib/images/votepage.png"
    import clickCandidate from "$lib/images/click-on-candidate.png"
    import voteForm from "$lib/images/vote-form.png"
</script>

<Hero>
    <div class="heroFormat">
        <h1 style="margin-top: 40px">
            How To Vote
        </h1>
        <p>
            Quick guide on how to vote
        </p>
    </div>
</Hero>

<div class="spaced-container">
    <h2>Register</h2>
    First you have to register
    <h4 style="text-align: center;">Click on the Register button on the home page</h4>
    <img src={register} alt="register button" style="margin-left: auto; margin-right: auto;max-width: 250px;"/>
    <h4 style="text-align: center;">Fill out the registration form and click Register</h4>
    <img src={registerForm} alt="register form" style="margin-left: auto; margin-right: auto;max-width: 350px;"/>
</div>
<div style="margin-top: 10px; height: 2px; width: 100%; background-color: var(--outline);"/>
<div class="spaced-container">
    <h2>Getting you account details</h2>
    To get your account details, wait for an email from us!
</div>
<div style="margin-top: 30px; height: 2px; width: 100%; background-color: var(--outline);"/>
<div class="spaced-container">
    <h2>Login</h2>
    Then you have to login to your account
    <h4 style="text-align: center;">Click on the Login button, in the top right, to go to the login page</h4>
    <img src={login} alt="login button" style="margin-left: auto; margin-right: auto;max-width: 250px;"/>
    <h4 style="text-align: center;">Fill out the login form and click Login</h4>
    <img src={loginForm} alt="login form" style="margin-left: auto; margin-right: auto;max-width: 350px;"/>
</div>
<div style="margin-top: 10px; height: 2px; width: 100%; background-color: var(--outline);"/>
<div class="spaced-container">
    <h2>Vote</h2>
    Now you are logged in, you have to vote
    <h4 style="text-align: center;">Click on the Vote button on the home page</h4>
    <img src={vote} alt="vote button" style="margin-left: auto; margin-right: auto;max-width: 250px;"/>
    <h4 style="text-align: center;">This is the vote page</h4>
    <img src={votePage} alt="vote page" style="margin-left: auto; margin-right: auto;max-width: 350px;"/>
    <h4 style="text-align: center;">Click on the candidate you want to vote for</h4>
    <img src={clickCandidate} alt="candidate button" style="margin-left: auto; margin-right: auto;max-width: 250px;"/>
    <h4 style="text-align: center;">Fill out the vote form and click Vote</h4>
    <img src={voteForm} alt="vote form" style="margin-left: auto; margin-right: auto;max-width: 350px;"/>
    <h3 style="text-align: center;">Success, You have voted!</h3>
</div>