module moduleA {
    // var with no value
    let varA: int;

    // constants
    const value = 125.3;
    const text: string = "testtando aaa";

    let text = "boa" - "asa";

    on "Start" (msg) {
        message msg1 = {
            "property": "teste"
        };
        share text ({});
        share "Test" (msg1);
    }

    // expressions
    let varB = ((12 + 3) -5 * 7 + value) > 35 && text || true;

    let a = b;

}