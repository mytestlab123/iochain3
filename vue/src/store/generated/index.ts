// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Mytestlab123Iochain3Iochain3 from './mytestlab123.iochain3.iochain3'


export default { 
  Mytestlab123Iochain3Iochain3: load(Mytestlab123Iochain3Iochain3, 'mytestlab123.iochain3.iochain3'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}