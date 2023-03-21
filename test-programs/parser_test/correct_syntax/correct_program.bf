module Test {
    // var with no value
    let varA: int;

    // constants
    const value = 125.3;
    const text: string = "testtando aaa";

    on "Start" (msg) {}

    // expressions
    let varB = ((12 + 3) -5 * 7 + value) > 35 && text || true;

    // array
    let arr: int[2][3][5];
    let arr2 = [2, 4, 6, 2, 4];

    //map
    let map: [string:float];
    let map1 = [string: [int: float[6]] ] {
        "teste": [int: float[6]] {
            2: 45.7,
            5: 3
        },
        "batata": [int: float[6]] {
            3: 57.2,
            5: 9
        }
    };

}