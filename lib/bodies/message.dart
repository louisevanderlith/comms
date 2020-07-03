class Message {
  final String message;
  final String email;
  final String name;
  final String phone;

  Message(this.message, this.email, this.name, this.phone);

  Map<String, dynamic> toJson() {
    return {
      "Body": message,
      "Email": email,
      "Name": name,
      "Phone": phone,
    };
  }
}
