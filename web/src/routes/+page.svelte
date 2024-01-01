<script>
	import { login } from "$lib/requests";

    let email = "";
    let password = "";
    let error = "";

    const performLogin = async () => {
        const loginResponse = await login(email, password);
        error = loginResponse?.error || "";
    }
</script>

<div>
    E-Mail
    <input bind:value={email} />
</div>
<div>
    Passwort
    <input bind:value={password} />
</div>
<div>
    <button on:click={performLogin}>Einloggen</button>
</div>
<div>
    {error}
</div>
