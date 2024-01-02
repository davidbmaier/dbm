<script>
    import { Input, Button } from 'flowbite-svelte'
    import { _ } from 'svelte-i18n';
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
    {$_("login.email")}
    <Input bind:value={email} />
</div>
<div>
    {$_("login.password")}
    <Input bind:value={password} />
</div>
<div>
    <Button on:click={performLogin}>{$_("login.login")}</Button>
</div>
<div>
    {error}
</div>
