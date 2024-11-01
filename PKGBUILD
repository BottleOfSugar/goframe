# Maintainer: Twoje Imię <twojemail@example.com>
pkgname=shell-star-ink
pkgver=1.0
pkgrel=1
pkgdesc="Powłoka podobna do Bash o nazwie Star Ink"
arch=('x86_64')
url="https://github.com/BottleOfSugar/goframe"
license=('MIT')
depends=('glibc')
makedepends=('go')
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz")

build() {
    cd "$srcdir/${pkgname}-${pkgver}/pkg"
    go build -o shell
}

package() {
    install -Dm755 "$srcdir/${pkgname}-${pkgver}/pkg/shell" "$pkgdir/usr/bin/shell-star-ink"
}

# Informacje opcjonalne
provides=('shell-star-ink')
conflicts=('shell-star-ink')
