package message

import "errors"

var (
	// 🔐 Authentication & Authorization
	ErrUnauthorized = errors.New("akses tidak diizinkan")
	ErrForbidden    = errors.New("akses ditolak")
	ErrInvalidToken = errors.New("token tidak valid")
	ErrTokenExpired = errors.New("token telah kedaluwarsa")
	ErrLoginFailed  = errors.New("email atau password salah")

	// 🔎 Resource & Data
	ErrNotFound      = errors.New("data tidak ditemukan")
	ErrAlreadyExists = errors.New("data sudah tersedia")
	ErrInvalidID     = errors.New("ID tidak valid")
	ErrEmptyResult   = errors.New("data kosong")

	// 📦 Request & Validation
	ErrBadRequest       = errors.New("permintaan tidak valid")
	ErrValidationFailed = errors.New("validasi gagal")
	ErrMissingField     = errors.New("field wajib belum diisi")
	ErrInvalidFormat    = errors.New("format data tidak sesuai")

	// ⚙️ Server & Internal
	ErrInternalServer  = errors.New("terjadi kesalahan pada server")
	ErrDB              = errors.New("gagal mengakses database")
	ErrCache           = errors.New("gagal mengakses cache")
	ErrFileIO          = errors.New("gagal membaca/menulis file")
	ErrExternalService = errors.New("gagal mengakses layanan eksternal")

	// 📦 Upload/File
	ErrFileTooLarge    = errors.New("ukuran file terlalu besar")
	ErrUnsupportedType = errors.New("tipe file tidak didukung")
	ErrUploadFailed    = errors.New("gagal mengunggah file")
)
