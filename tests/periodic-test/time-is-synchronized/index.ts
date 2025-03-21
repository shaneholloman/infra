import { Sandbox } from "npm:@e2b/code-interpreter";


// Helper function to stream command output
async function streamCommandOutput(command: string, args: string[]) {
    const cmd = new Deno.Command(command, {
        args: args,
        stdout: "piped",
        stderr: "piped",
    });

    const process = cmd.spawn();
    const decoder = new TextDecoder();

    let output = ''

    // Stream stdout
    for await (const chunk of process.stdout) {
        console.log(decoder.decode(chunk));
        output += decoder.decode(chunk)
    }

    // Stream stderr
    for await (const chunk of process.stderr) {
        console.error(decoder.decode(chunk));
        output += decoder.decode(chunk)
    }

    // Wait for the process to complete and get the status
    const status = await process.status;
    return { status, output }
}

const uniqueID = crypto.randomUUID();
const templateName = `test-template-${uniqueID}`
console.log('ℹ️ templateName:', templateName)

// Build template command with streaming output
console.log(`Building template ${templateName}...`);
const buildStatus = await streamCommandOutput('npx', [
    '@e2b/cli',
    'template',
    'build',
    '--name',
    templateName,
]);

if (buildStatus.status.code !== 0) {
    throw new Error(`❌ Build failed with code ${buildStatus.status.code}`);
}

console.log('✅ Template built successfully')

// read template id from e2b.toml
const e2bToml = await Deno.readTextFile('e2b.toml')
const templateID = e2bToml.match(/template_id = "(.*)"/)?.[1]

if (!templateID) {
    throw new Error('❌ Template ID not found in e2b.toml')
}

// sleep for 5 seconds to create a time delta
await new Promise(resolve => setTimeout(resolve, 5000))



try {

    // remove the file to make script idempotent in local testing
    await Deno.remove('e2b.toml')

    if (!templateID) {
        throw new Error('❌ Template not found')
    }
    console.log('ℹ️ creating sandbox')
    const sandbox = await Sandbox.create(templateID, { timeoutMs: 10000 })
    console.log('ℹ️ sandbox created')


    console.log('ℹ️ running command')

    console.log('ℹ️ starting command')
    const localDate = new Date().getTime() / 1000
    const date = await sandbox.commands.run('date +%s')
    console.log('localDate', localDate)

    console.log('date', date.stdout)
    const dateUnix = parseInt(date.stdout)
    console.log('ℹ️ comparing dates', dateUnix, localDate)

    // compare the dates, should be within 1 second
    if (Math.abs(dateUnix - localDate) > 1000) {
        throw new Error('❌ Date is not synchronized')
    }

    console.log('✅ date is synchronized')

    // kill sandbox
    await sandbox.kill()

} finally {
    // delete template
    const output = await streamCommandOutput('npx', [
        '@e2b/cli',
        'template',
        'delete',
        '-y',
        templateID
    ])

    if (output.status.code !== 0) {
        throw new Error(`❌ Delete failed with code ${output.status.code}`);
    }

}