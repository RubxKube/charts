message("Please provide a clear title.") unless pr.title.match?(/^\[.+?\]/)
warn("This PR does not include a description.") if pr.body.empty?

if github.pr_body.length < 5
    fail "Please provide a summary in the Pull Request description"
end

lgtm.check_lgtm