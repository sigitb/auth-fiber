package templates

var header = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css?family=Nunito:400,600,700,800,900&display=swap" rel="stylesheet">
    <link
        href="https://fonts.googleapis.com/css2?family=Rubik:ital,wght@0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap"
        rel="stylesheet">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css"
        integrity="sha384-B0vP5xmATw1+K9KRQjQERJvTumQW0nPEzvF6L/Z6nronJ3oUOFUFpCjEUQouq2+l" crossorigin="anonymous">
</head>

<body style="font-family: Nunito, sans-serif;">
    <div style=" width: 100%; height: 200px; border-radius: 0 0 100% 100%; background-color: #007cff;">
    </div>
	<div style="padding: 0px 10px 0px 10px">
		<div style="display: flex;justify-content: center;">
			<div
				style="width:100%;border: 2px solid #D7E1EA; border-radius:10px; padding:10px 30px 10px 30px; margin-top:20px">

`

var EmailFotter = `<div>
<p style="font-weight: 500; font-size: 14px; color: #1f1f1f;">Terima Kasih,
	<br>
	Base AUTH
</p>
</div>
<div style="background-color: #D7E1EA;height:2px"></div>
<div style="display:flex;margin-top:20px">
<div style="margin-left: 20px">
	<div style="margin-top: 5px">
		<b style="font-weight: 700;font-size: 20px;color: #1f1f1f;"><strong>Base AUTH</strong></b>
	</div>
	<div style="margin-top: 5px">
		<b style="  font-size: 16px;margin-top:10px;
		font-weight: 500;
		color: #6C6C6C">Jl. Oke Siap no 102</b>
	</div>
	<div>
		<b style="  font-size: 16px;margin-top:10px;
		font-weight: 500;
		color: #6C6C6C">(123) 456789</b>
	</div>
	<div>
		<b style="  font-size: 16px;margin-top:10px;
		font-weight: 500;
		color: #6C6C6C">baseauth@gmail.com</b>
	</div>

</div>
</div>
</div>
</div>
</div>
</body>

</html>
`

func TemplateResetPassword(url string, name string) string {

	data := header + `
	<p>
	Hai, <span class="font-weight-bold" style="font-weight: bold;">` + name + `</span>
	<br>
	<p style="margin-bottom: 30px; font-weight: 500; font-size: 14px; color: #1f1f1f;">
		Sepertinya anda melakukan permintaan untuk mengatur ulang password. Jika benar, silahkan klik tombol berikut untuk mengatur ulang password anda.
	</p>
	<br>
	<center>
		<a href="` + url + `" style="background: #215480; padding: 15px; border-radius: 30px; font-weight: 600; font-size: 14px; color: #FFFFFF; text-decoration: none;">Atur Ulang Password</a>
	</center>
	<br>
	<p style="font-weight: 500; font-size: 14px; color: #1f1f1f;">
		Jika Anda tidak merasa melakukan pengaturan ulang atau tidak ingin perubahan password, silahkan abaikan e-mail ini.
	<p>
	<br>
	<p> jika ada masalah dengan button di atas copy link berikut <p>
	<p style="font-weight: 500; font-size: 14px; color: #1f1f1f;">
	<a href="` + url + `">"` + url + `"</a>
	<p>
	</p>
	` + EmailFotter

	return data
}

func TempalteVerification(name string, email string, kode string, url string) string {

	data := header + `
	<br>
		<p>
		Hai, <span class="font-weight-bold" style="font-weight: bold;">` + name + `</span>
		<br>
		Akun anda telah terdaftar di Base Auth ini dengan data sebagai berikut
		<br>
		<br>
		Nama : ` + name + `
		<br>
		E-mail : ` + email + `
		<br>
		<br>
		Dan gunakan Kode Verifikasi berikut untuk mengaktifkan akun Anda:
		<p style="font-weight: 700;font-family: 'Rubik', sans-serif;font-size: 32px;text-align: center;">` + kode + `</p>
		Jika Anda tidak merasa melakukan verifikasi e-mail, segera hubungi Kami melalui baseauth@gmail.com
		</p>
	` + EmailFotter
	return data
}