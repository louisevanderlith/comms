class Message {
  final String name;
  final String subject;
  final String email;
  final String phone;
  final String body;
  final String to;

  Message(this.name, this.subject, this.email, this.phone, this.body, this.to);

  Map<String, dynamic> toJson() {
    return {
      "Name": name,
      "Subject":subject,
      "Email": email,
      "Phone": phone,
      "Body": body,
      "To": to,
    };
  }
}
