package cronet

// #cgo LDFLAGS: -Wl,--build-id=sha1 -fPIC -Wl,-z,noexecstack -Wl,-z,relro -Wl,-z,now -Wl,-z,max-page-size=4096 -Wl,--icf=all -Wl,--color-diagnostics -Wl,-mllvm,-instcombine-lower-dbg-declare=0 -flto=thin -Wl,--thinlto-jobs=all -Wl,--thinlto-cache-dir=thinlto-cache -Wl,--thinlto-cache-policy=cache_size=10%:cache_size_bytes=40g:cache_size_files=100000 -Wl,-mllvm,-import-instr-limit=5 -fwhole-program-vtables -no-canonical-prefixes -rdynamic -Wl,-z,defs -Wl,--as-needed -Wl,--lto-O2 -pie -Wl,--disable-new-dtags
import "C"
