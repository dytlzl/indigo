setTimeout(async () => {
    await navigator.clipboard.writeText(
        JSON.stringify(
            Array.from(
                document.getElementsByClassName('plan-option')
            ).map(
                e => {
                    args = JSON.parse(
                        "[" +
                        e.outerHTML
                            .replace(/\n/g, '')
                            .replace(/'/g, '"')
                            .replace(/.*chooseInstaneType\(([^)]*)\).*/, '$1') +
                        "]"
                    )
                    code = args[2]
                    parts = code.replace(/([0-9]+)vCPU([0-9]+)([MG])B([0-9]+)GB/, '$1,$2,$3,$4').split(',')
                    vCpu = Number(parts[0])
                    ram = Number(parts[2] === 'M' ? parts[1] : parts[1] * 1024)
                    ssd = Number(parts[3])
                    return {
                        id: Number(e.classList[4].split('-')[3]),
                        // description: e.innerText.replace(/\n/g, ', '),
                        code: code,
                        ipType: args[3],
                        vCpu: vCpu,
                        ram: ram,
                        ssd: ssd,
                        network: e.innerText.split('\n').slice(-1)[0],
                    }
                }
            ),
            null,
            2
        )
    ).then(() => console.log('copied to clipboard'), (reason) => console.log(reason))
}, 1000)