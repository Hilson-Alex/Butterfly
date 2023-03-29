module ModuleB {

    // array
    let arr: int[2][3][5];

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

    on "Test" (msg) {
        let arr2 = [2, 4, 6, 2, 4];
        let b = msg.property;
        for (let a = 0; i < 5; i++) {
            arr[i] += 2 * i;
        }
        while (true) {
            let a = 0;
        }
        do {
            map = map1;
        } while (0 < 10);
    }

}