Instalasi Go dan pengaturan lingkungan (GOPATH, GOROOT).

1. Instalasi golang : 
-Download terlebih dahulu installer-nya di https://golang.org/dl/. Pilih installer untuk sistem operasi Windows sesuai jenis bit yang digunakan.

-Setelah ter-download, jalankan installer, klik next hingga proses instalasi selesai. By default jika anda tidak merubah path pada saat instalasi, Go akan ter-install di C:\go. Path tersebut secara otomatis akan didaftarkan dalam PATH environment variable.

-Buka Command Prompt / CMD, eksekusi perintah "go version" untuk mengecek versi Go.

Jika output adalah sama dengan versi Go yang ter-install, menandakan proses instalasi berhasil.

2. setting environment :
-GOROOT
secara default go akan terinstal di c:\go disitu akan secara otomatis muncul environment variable GOROOT untuk mengecek nya bisa akses di cmd deengan menuliskan perintah go env
setelah itu maka akan muncul keterangan environment yang ada di golang. disitu akan ada tulisan GOROOT dan ketika di klik akan menampilkan root untuk folder go di device kalian

-GOPATH
untuk GOPATH bisa diubah sesuai tempat folder yang kita inginkan, cara merubah folder gopath itu ialah kita tinggal masuk ke control panel lalu cari view advance system settings setelah itu
klik environment variable dan pilih new dan masukan GOOPATH dan folder yang ingin dibuatnya.
