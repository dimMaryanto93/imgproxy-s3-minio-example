<?php

$key = '6e45924d20fef273429090a206eb737c86ce1c7dccc8275cb313fc8a0d1ddcf38481535ad556c41a566a5c3eebda84c44271a3f2ee1bee48728689a307469c15';
$salt = '182ef2016fc483dcac88666ae9813469b94707f903d8f83efdc2f22ebbf9c5521c20330d9a6c326054c3524d5e08f8ba2e1a957dbdfc7777a4b1843e21fd2d14';

$keyBin = pack("H*" , $key);
if(empty($keyBin)) {
	die('Key expected to be hex-encoded string');
}

$saltBin = pack("H*" , $salt);
if(empty($saltBin)) {
	die('Salt expected to be hex-encoded string');
}

$encodedSourceImage = "http%3A%2F%2F192.168.100.251%3A39000%2Fdigivetmr%2F0a9349e5-c1b1-4d7e-b964-fc9de0587ac9.JPG%3FX-Amz-Algorithm%3DAWS4-HMAC-SHA256%26X-Amz-Credential%3DK24UQM1DG1U008IQFFGO%252F20221108%252Fus-east-1%252Fs3%252Faws4_request%26X-Amz-Date%3D20221108T033358Z%26X-Amz-Expires%3D604800%26X-Amz-Security-Token%3DeyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJLMjRVUU0xREcxVTAwOElRRkZHTyIsImV4cCI6MTY2Nzg4MjAxNSwicGFyZW50IjoidGFiZWxkYXRhIn0.g-aAtpaqb0-RkEr0tMixw7u2d5aaefRI14BuP1CiZZpBwLaaI-XPJEmbxtqsesubhi9ByhyVY_68Jj9aozxlfw%26X-Amz-SignedHeaders%3Dhost%26versionId%3Dnull%26X-Amz-Signature%3D8468fcd4661875bb2a9815b6bb3606fd28fb00b8f616c3bb6ec354916d917d93";
$path = "/rs:fit:300:300/plain/".$encodedSourceImage;

$signature = rtrim(strtr(base64_encode(hash_hmac('sha256', $saltBin.$path, $keyBin, true)), '+/', '-_'), '=');

print(sprintf("/%s%s", $signature, $path));