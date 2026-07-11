import { render } from "../utils/render.js";
import { appLayout } from "../layouts/appLayout.js";
import { postDetailsTemplate } from "../templates/postTemplate.js";
import { initNavbar } from "../components/navbar.js";
import { getPost } from "../services/postService.js";
import { getComments } from "../services/commentService.js";
import { initCommentForm } from "../components/commentForm.js";
export async function postDetailsPage(id) {

    try {

        const post = await getPost(id);

        const comments = await getComments(id);
        
       
        console.log(comments);

        render(
            appLayout(
                postDetailsTemplate(post, comments)
            )
        );

        initNavbar();

        initCommentForm(post.id);

    } catch (error) {

        render(
            appLayout(`
                <h2>
                    ${error.message}
                </h2>
            `)
        );

    }

}

