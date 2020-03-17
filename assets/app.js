// Delete the notification
document.getElementById("notification-delete").onclick = function() {
    document.getElementById("notification-modal").className = "modal"
}

// Read the content of a file and put it into the number field
document.getElementById("read").onclick = function() {
    const input = document.querySelector("input[type=file]")
    if (input.files.length === 0) {
        console.log('No file selected.')
        return
    }

    const reader = new FileReader();
    reader.onload = function fileReadCompleted() {
        let content = reader.result.substring(
            0,
            parseInt(document.getElementById("digits").value)
        )

        document.getElementById("number").value = content
    };
    reader.readAsText(input.files[0]);
}