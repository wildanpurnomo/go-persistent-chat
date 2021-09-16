export default {
  data: () => ({
    usernameRules: [
      (v) => !!v || "Mohon masukkan username",
      (v) => (v || "").length >= 6 || "Minimal 6 karakter",
    ],
    passwordRules: [
      (v) => (v || "").length >= 8 || "Minimal 8 karakter",
      (v) => /(?=.*[a-z])/.test(v) || "Harus terdapat minimal satu huruf kecil",
      (v) => /(?=.*[A-Z])/.test(v) || "Harus terdapat minimal satu huruf besar",
      (v) => /(?=.*[0-9])/.test(v) || "Harus terdapat minimal satu digit angka",
    ],
    isFormValid: false,
    isFormLoading: false,
    isPasswordShown: false,
    errorMessage: "",
  }),
};
