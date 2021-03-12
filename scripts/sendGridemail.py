import argparse
from sendgrid import SendGridAPIClient
from sendgrid.helpers.mail import Mail, Cc
import codecs

parser = argparse.ArgumentParser()
parser.add_argument("contentfile", help="email content")
parser.add_argument("email", help="mail list")
parser.add_argument("subject", help="mail subject")
# parser.add_argument("release_version", help="release version")
parser.add_argument("api_key", help="sendgrid api key")
parser.add_argument("cc", help="cc mail list")
args = parser.parse_args()

from_email = 'codegen@azure-devex-tools.com'
to_emails = args.email.split(';')
subject = args.subject

cc = args.cc.split(';')
cc_emails = []
for c in cc:
  if len(c) > 0 :
    cc_emails.append(Cc(c, c))

f = codecs.open(args.contentfile, 'r')
html_content = f.read()
print(html_content)
# subject = 'Onboading ' + args.release_version + ' is Released'
# f = codecs.open(args.template, 'r')
# html_content = f.read()
# html_content = html_content.replace("trenton_version", args.release_version)

message = Mail(
    from_email=from_email,
    to_emails=to_emails,
    subject=subject,
    html_content=html_content)

if len(cc_emails) > 0:
    message.add_cc(cc_emails)

try:
    sg = SendGridAPIClient(api_key=args.api_key)
    response = sg.send(message)
    print(response.status_code)
    print(response.body)
    print(response.headers)
except Exception as e:
    print(e.message)