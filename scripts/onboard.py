from jinja2 import Environment, FileSystemLoader
import sys

# file_loader = FileSystemLoader('./')
# env = Environment(loader=file_loader)

# template = env.get_template('script/onboardTemplate')
# users = [{
#     "url": "url",
#     "username": "test"
# }]
# output = template.render(users)
# print(output)

def Generate_HTML(title, head, codePR, triggerPR, submitURL, customizeURL):
  file_loader = FileSystemLoader('./')
  env = Environment(loader=file_loader)
  template = env.get_template('scripts/onboardTemplate')
  output = template.render(title=title, head=head, codePR=codePR, triggerPR=triggerPR, submitURL=submitURL, customizeURL=customizeURL);
  print(output)
  return output


if __name__ == "__main__":
#   print(f"Arguments count: {len(sys.argv)}")
  if sys.argv[1] == "--help":
    print("usage: python scripts/emailsender.py [title] [head] [codePR] [triggerPR] [submitURL] [customizeURL]")
    
#   for i, arg in enumerate(sys.argv):
#     print(f"Argument {i:>6}: {arg}")
  title = sys.argv[1]
  head = sys.argv[2]
  codePR = sys.argv[3]
  triggerPR = sys.argv[4]
  submitURL = sys.argv[5]

  customizeURL = sys.argv[6]

  Generate_HTML(title, head, codePR, triggerPR, submitURL, customizeURL)